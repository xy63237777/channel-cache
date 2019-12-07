package gocache

import (
	"fmt"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	cache := NewSignalCache(CacheForLRU, 2)
	cache.Set("hello", 1)
	cache.Set("aaaa", 2)
	cache.Set("cccc", 3)
	fmt.Println(cache.Get("hello"))
	fmt.Println(cache.Get("aaaa"))
	fmt.Println(cache.Get("cccc"))
	cache.Set("kmp","3333")
	fmt.Println(cache.Get("aaaa" + ""))
}
