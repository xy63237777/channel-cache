package gocache

import (
	"errors"
	"log"
	"sync"
	"time"
)

/**
*****************************************************
*				   Signal Cache                     *
*****************************************************
**/

var cFlag = false
var onceSignalCache *Cache
var signalMu sync.Mutex
var onceSignal sync.Once
func NewSignalCache(typeCache CacheType, capacity int) *Cache {

	return newSignalCacheOnce(typeCache, capacity)
}

func NewSignalCacheCustom(ca CacheInter) *Cache {
	return newCacheForCustom(ca)
}

func NewSignalCacheForDefault() *Cache {
	return newSignalCacheOnce(CacheForLFU, DefaultCapacity)
}

func newSignalCacheOnce(typeCache CacheType, capacity int) *Cache {
	ch := make(chan *Cache,1)
	if cFlag {
		return onceSignalCache
	}
	onceSignal.Do(func() {
		if LoggerLevel == DEBUG {
			log.Println("NewSignalCache is once call...")
		}
		onceSignalCache = newCache(typeCache, capacity)
		cFlag = true
		ch <- onceSignalCache
	})
	select {
	case cache := <- ch: return cache
	case _ = <- time.After(time.Second):
		log.Println("NewSignalCache call timeout...")
		return onceSignalCache
	}
}


func StopSignalCache() error {
	if cFlag {
		return errors.New("signal Cache not start")
	}
	signalMu.Lock()
	defer signalMu.Unlock()
	return onceSignalCache.close("signal")
}

func RunSignalCache() error {
	if cFlag {
		return errors.New("signal Cache not start")
	}
	signalMu.Lock()
	defer signalMu.Unlock()
	return onceSignalCache.run("signal")
}
