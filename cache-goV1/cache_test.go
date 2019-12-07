package gocache

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCache_GetAsync(t *testing.T) {
	cache := NewSignalCacheForDefault()
	cache.Set("hello", "world")
	fmt.Println(cache.Get("hello"))
	async := cache.GetAsync("hello")
	fmt.Println(<- async)
	cache.SetForExpiration("kkm", "world",time.Second)
	time.Sleep(time.Second)
	fmt.Println(cache.Get("kkm"))
	fmt.Println(<- cache.GetAsync("kkm"))
	//fmt.Println(cache.SetForExpiration("hello","pppp",time.Second))
	fmt.Println(cache.Get("hello"))
	fmt.Println(<- cache.GetAsync("hello"))
	time.Sleep(time.Second)
	fmt.Println(cache.Get("hello"))
	fmt.Println(<- cache.GetAsync("hello"))

}


func TestCache_Get(t *testing.T) {
	n := 500000
	cache := NewSignalCacheForDefault()
	start := time.Now().UnixNano()
	for i:=0; i < n; i++ {
		k := i % 4096
		cache.Set(strconv.Itoa(k), "hello")
		cache.GetAsync(strconv.Itoa(k))
	}
	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
	start = time.Now().UnixNano()
	for i:=0; i < n; i++ {
		//cache.SetFor(strconv.Itoa(i), "hello")
		//cache.GetFor(strconv.Itoa(i))
	}
	fmt.Println((time.Now().UnixNano() - start) / int64(time.Millisecond))
}

func hello()  {

}

func TestFunc(t *testing.T)  {
	cache := NewSignalCacheForDefault()
	cache.Set("hello", "world")
	cache.Set("hello", "world")
	cache.Set("hello", "world")

	fn := hello
	p := &fn
	ch := make(chan interface{}, 1)
	fmt.Printf("%p\n", p)
	ch <- p
	pp := <-ch
	km := pp.(*func())
	(*km)()
	fmt.Printf("%p\n", pp)
	fmt.Printf("%p\n", km)
}