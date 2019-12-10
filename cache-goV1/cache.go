package gocache

import (
	"errors"
	"fmt"
	"four-seasons/log"
	"time"
)

func newCacheForCustom(ca CacheInter) *Cache {
	cache := newCache("", 0)
	cache.cache = ca
	return cache
}

func newCache(typeCache CacheType, capacity int) *Cache {
	var cache *Cache
	if capacity <= 0 {
		capacity = DefaultCapacity
	}
	cache =  &Cache{cache:nil,
		dp:newDefaultDispatcher(),
		expiration:DefaultExpiration,
		liquidator:newLiquidator(),
		onceOceChannel:make(chan interface{}, 1),
	}
	if typeCache == CacheForLRU {
		cache.cache = newLRUCache(capacity, cache)
	} else if typeCache == CacheForEasy {
		cache.cache = newEasyCache(capacity)
	} else if typeCache == "" {
		//nothing
	} else {
		cache.cache = newLFUCache(capacity, cache)
	}

	cache.dp.start(cache.liquidator)
	return cache
}


func newCacheForDefault() *Cache {
	return newCache(CacheForLFU, DefaultCapacity)
}


/*
-----------------------------------------------------------------------------

*/

func (c *Cache) IsClose() bool {
	return c.isClose
}


func (c *Cache) close(name string) error {
	if c.isClose {
		log.Error("Cache name : %s is Closed", name)
		return errors.New("Cache Name : " + name +" is Closed")
	}
	log.Info("Closing cache Name is %s ...",name)
	c.isClose = true
	for i := 0; i < len(c.dp.queue); i++ {
		c.dp.stateCh <- CLOSE
	}
	return nil
}

func (c *Cache) run(name string) error {
	if !c.isClose {
		log.Error("Cache name : %s is Running", name)
		return errors.New("Cache Name : " + name +" is Running")
	}
	log.Info("Start cache Name is %s ...",name)
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


func (c *Cache) Set(key string, val interface{})  {
	if checkCloseGentle(c) {
		return
	}
	c.setAsync(key, newItem(val, NoExpiration))
}

func (c *Cache) SetForDefaultExpiration(key string, val interface{})  {
	if checkCloseGentle(c) {
		return
	}
	doPush(c, newClearNode(c, key, c.expiration))
	c.setAsync(key, newItem(val, c.expiration))
}

func doPush(c *Cache, node *needClearNode)  {
	c.liquidator.push(node)
	if c.liquidator.elements.Length() >=  c.cache.Capacity() {
		 c.liquidator.clearFunc()
	}

}

func (c *Cache) SetForExpiration(key string, val interface{}, d time.Duration) {
	if checkCloseGentle(c) {
		return
	}
	doPush(c, newClearNode(c, key, d))
	c.setAsync(key, newItem(val, d))
}


/*
-----------------------------------------------------------------------------

*/

func (c *Cache) setAsync(key string, val *item)  {
	setFunc(key, val, c)
}



var onceSetFunc = func(data *commonsData)  {
	data.master.cache.Put(*data.key, data.val)
}

func setFunc(key string, val *item ,c *Cache)  {
	nc :=newCommons(&onceSetFunc, newCommonsData(key, val, c, nil))
	c.dp.queue <- nc
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

var onceGetFunc = func(data *commonsData) {
	obj := data.master.cache.Get(*data.key)
	if obj == nil {
		data.out <- nil
	} else {
		item := obj.(*item)
		if  item.IsExpired() {
			data.out <- nil
		} else {
			data.out <- item.obj
		}
	}
}



func getHash(key *string, size int) byte {
	hash := (*key)[0]%uint8(size)
	fmt.Println(hash, *key)
	return hash
}

func getFunc(key string, c *Cache, ch chan interface{})  {
	c.dp.queue <- newCommons(&onceGetFunc, newCommonsData(key, nil, c,ch))
}




func (c *Cache) get(key string) (interface{} ,bool) {
	getFunc(key, c, c.onceOceChannel)
	obj := <- c.onceOceChannel
	if obj == nil {
		return obj, false
	}
	return obj, true
}

func (c *Cache) getAsync(key string) chan interface{} {
	ch := make(chan interface{}, 1)
	getFunc(key, c, ch)
	return ch
}



/*
-----------------------------------------------------------------------------
*/

func (c *Cache) Delete(key string) {
	if checkCloseGentle(c) {
		return
	}
	c.delAsync(key)
}

//func (c *Cache) Delete(key string) interface{} {
//	if checkCloseGentle(c) {
//		return nil
//	}
//	return c.del(key)
//}

//func (c *Cache) DeleteAsync(key string) interface{} {
//	if checkCloseGentle(c) {
//		return nil
//	}
//	return c.delAsync(key)
//}

//func (c *Cache) del(key string) interface{} {
//	ch := make(chan interface{}, 1)
//	delFunc(key, c, ch)
//	return <- ch
//}


func (c *Cache) delAsync(key string)  {
	delFunc(key, c)
}

//func (c *Cache) delAsync(key string) chan interface{} {
//	ch := make(chan interface{}, 1)
//	delFunc(key, c, ch)
//	return ch
//}

var onceDelFunc = func(data *commonsData) {
	data.master.cache.Delete(*data.key)
}

func delFunc(key string ,c *Cache)  {
	c.dp.queue <- newCommons(&onceDelFunc, newCommonsData(key, nil, c, nil))
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


