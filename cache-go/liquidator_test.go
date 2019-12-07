package gocache

import (
	"fmt"
	"testing"
	"time"
)

func TestLiquidator(t *testing.T) {
	cache := NewSignalCacheForDefault()
	cache.SetForExpiration("hello", "world", time.Second)
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 300)
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 1200)
	fmt.Println(cache.Get("hello"))

	fmt.Println(<- cache.GetAsync("hello"))
}
