package gocache

import (
	"errors"
	"github.com/gpmgo/gopm/modules/log"
	"time"
)

func newCache(typeCache CacheType, capacity int) *Cache {
	var cache *Cache
	if capacity <= 0 {
		capacity = DefaultCapacity
	}
	if typeCache == CacheForLRU {
		cache =  &Cache{cache:nil,
			dp:newDefaultDispatcher(),
			expiration:DefaultExpiration,
		}
		cache.cache = NewLRUCache(capacity, cache)
	} else {
		cache = newCacheForDefault()
	}
	if cache.liquidator == nil {
		cache.liquidator = newLiquidator()
	}
	cache.dp.start(cache.liquidator)
	return cache
}


func newCacheForDefault() *Cache {
	cache := &Cache{
		dp:newDefaultDispatcher(),
		expiration:DefaultExpiration,
		liquidator:newLiquidator(),
	}
	cache.cache = NewLFUCache(DefaultCapacity, cache)
	return cache
}


/*
-----------------------------------------------------------------------------

*/

func (c *Cache) IsClose() bool {
	return c.isClose
}


func (c *Cache) close(name string) error {
	if c.isClose {
		return errors.New("Cache Name : " + name +" is Closed")
	}
	c.isClose = true
	c.dp.stateCh <- CLOSE
	return nil
}

func (c *Cache) run(name string) error {
	if !c.isClose {
		return errors.New("Cache Name : " + name +" is Running")
	}
	c.isClose = false
	c.dp.start(c.liquidator)
	return nil
}

/*
-----------------------------------------------------------------------------

*/
func (c Cache) GetDefaultExpiration() time.Duration {
	return c.expiration
}

func (c *Cache) SetDefaultExpiration(newExpiration time.Duration) {
	c.expiration = newExpiration
}

/*
-----------------------------------------------------------------------------

*/

func (c *Cache) Set(key string, val interface{}) interface{} {
	if checkCloseGentle(c) {
		return nil
	}
	return c.set(key, newItem(val, NoExpiration))
}


func (c *Cache) SetForDefaultExpiration(key string, val interface{}) interface{} {
	if checkCloseGentle(c) {
		return nil
	}
	c.liquidator.elements.Push(newClearNode(c, key, c.expiration))
	return c.set(key, newItem(val, c.expiration))
}

func (c *Cache) SetForExpiration(key string, val interface{}, d time.Duration) interface{} {
	if checkCloseGentle(c) {
		return nil
	}
	c.liquidator.elements.Push(newClearNode(c, key, d))
	return c.set(key, newItem(val, d))
}

func (c *Cache) SetForDefaultExpirationAsync(key string, val interface{} ) interface{} {
	if checkCloseGentle(c) {
		return nil
	}
	c.liquidator.elements.Push(newClearNode(c, key, c.expiration))
	return c.setAsync(key, newItem(val, c.expiration))
}

func (c *Cache) SetForExpirationAsync(key string, val interface{}, d time.Duration) interface{} {
	if checkCloseGentle(c) {
		return nil
	}
	c.liquidator.elements.Push(newClearNode(c, key, d))
	return c.setAsync(key, newItem(val, d))
}


func (c *Cache) SetAsync(key string, val interface{}) chan interface{} {
	if checkCloseGentle(c) {
		return nil
	}
	return c.setAsync(key, newItem(val, NoExpiration))
}

/*
-----------------------------------------------------------------------------

*/

func (c *Cache) set(key string, val *item) interface{} {
	ch := make(chan interface{}, 1)
	setFunc(key, val, c, ch)
	return <- ch
}

func (c *Cache) setAsync(key string, val *item) chan interface{} {
	ch := make(chan interface{}, 1)
	setFunc(key, val, c, ch)
	return ch
}

func setFunc(key string, val *item ,c *Cache, ch chan interface{})  {
	fn := func() {
		obj := c.cache.Put(key, val)
		if obj == nil {
			ch <- nil
		}else {
			item := obj.(*item)
			if  item.IsExpired() {
				ch <- nil
			} else {
				ch <- item.obj
			}
		}
	}

	c.dp.queue <- newCommons(fn)
}

/*
-----------------------------------------------------------------------------

*/



func (c *Cache) Get(key string) (interface{}, bool) {
	if checkCloseGentle(c) {
		return nil, false
	}
	return c.get(key)
}


func (c *Cache) GetAsync(key string) chan interface{} {
	if checkCloseGentle(c) {
		return nil
	}
	return c.getAsync(key)
}

func (c *Cache) get(key string) (interface{} ,bool) {
	ch := make(chan interface{}, 1)
	c.dp.queue <- newCommons(func() {
		ch <- c.cache.Get(key)
	})
	obj := <-ch
	if obj == nil{
		return nil, false
	} else {
		item := obj.(*item)
		if item.IsExpired() {
			return nil, false
		}
		return item.obj, true
	}
}

func (c *Cache) getAsync(key string) chan interface{} {
	ch := make(chan interface{}, 1)
	c.dp.queue <- newCommons(func() {
		obj := c.cache.Get(key)
		if obj == nil {
			ch <- nil
		} else {
			item := obj.(*item)
			if item.IsExpired() {
				ch <- nil
			} else {
				ch <- item.obj
			}
		}
	})
	return ch
}



/*
-----------------------------------------------------------------------------
*/

func (c *Cache) Delete(key string) interface{} {
	if checkCloseGentle(c) {
		return nil
	}
	return c.del(key)
}

func (c *Cache) DeleteAsync(key string) interface{} {
	if checkCloseGentle(c) {
		return nil
	}
	return c.delAsync(key)
}

func (c *Cache) del(key string) interface{} {
	ch := make(chan interface{}, 1)
	delFunc(key, c, ch)
	return <- ch
}


func (c *Cache) delAsync(key string) chan interface{} {
	ch := make(chan interface{}, 1)
	delFunc(key, c, ch)
	return ch
}

func delFunc(key string ,c *Cache, ch chan interface{})  {
	c.dp.queue <- newCommons(func() {
		obj := c.cache.Delete(key)
		if obj == nil {
			ch <- nil
		} else {
			item := obj.(*item)
			if item.IsExpired() {
				ch <- nil
			} else {
				ch <- item.obj
			}
		}
	})
}

/*
-----------------------------------------------------------------------------
*/

func checkClose(c *Cache)  {
	if c.isClose {
		panic("Cannot use a stopped Cache\n" +
			"It is not recommended that you use stop to stop\n" +
			"A better way is to use signal")
	}
}

func checkCloseGentle(c *Cache) bool {
	if c.isClose {
		log.Warn("Warning...  Using a stopped cache")
		return true
	}
	return false
}

