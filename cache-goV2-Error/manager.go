package gocache

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	DEBUG = iota
	DEFAULT
)


/**
得到缓存
 */

var mFlag = false
var onceCacheManager *CacheManager
var LoggerLevel = DEFAULT


func NewCacheManager() *CacheManager {
	return newCacheManagerOnce()
}

var onceManagerChanel = make(chan *CacheManager,1)

var once sync.Once

func newCacheManagerOnce() *CacheManager {
	if mFlag {
		return onceCacheManager
	}
	once.Do(func() {
		if LoggerLevel == DEBUG {
			log.Println("NewCacheManger is once call...")
		}

		onceCacheManager = newCacheManager()
		mFlag = true
		onceManagerChanel <- onceCacheManager
	})
	select {
	case manager := <- onceManagerChanel: return manager
	case _ = <- time.After(time.Second*5):
		log.Println("NewCacheManger call timeout...")
		return onceCacheManager
	}
}

func newCacheManager() *CacheManager {
	return &CacheManager{
		mu: sync.RWMutex{},
		m:  make(map[string]*Cache),
	}
}



func (cm *CacheManager) GetCache(key string) *Cache {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.m[key]
}

/**
根据规则创建缓存
*/
func (cm *CacheManager) CreateCache(key string,typeCache CacheType, capacity int) *Cache {
	fmt.Println("hello123123123")
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cache := newCache(typeCache, capacity)
	cm.m[key] = cache
	return cache
}

/**
创建一个缓存实例
*/
func (cm *CacheManager) CreateCacheForDefault(key string) *Cache {

	cm.mu.Lock()
	defer cm.mu.Unlock()
	cache := newCacheForDefault()
	cm.m[key] = cache
	return cache
}

/**
添加一个缓存实例
*/
func (cm *CacheManager) AddCache(key string, cache *Cache) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.m[key] = cache
}

/**
删除缓存并且返回给客户端
*/
func (cm *CacheManager) DeleteCache(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cache,ok := cm.m[key]
	if ok {
		delete(cm.m, key)
	}
	_ = cache.close(key)
}

func (cm *CacheManager) CacheRun(key string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cache,ok := cm.m[key]
	if !ok {
		return errors.New("Cache name : " + key + "not found")
	}
	return cache.run(key)
}

func (cm *CacheManager) CacheStop(key string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cache,ok := cm.m[key]
	if !ok {
		return errors.New("Cache name : " + key + "not found")
	}
	return cache.close(key)
}




