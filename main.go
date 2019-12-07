package main

import (
	"fmt"
	_ "sync/atomic"
)

type item struct {
	obj interface{}
	aa int64
}

func main() {
	//cache := cache_go.NewSignalCache(cache_go.CacheForLFU, 2)
	//cache.Set("hello","aaaa")
	//cache.Set("hello","ccccc")
	//cache.Set("cccccc", "tttt")
	//fmt.Println(cache.Get("hello"))
	//fmt.Println(cache.Get("bbbbb"))
	//fmt.Println(cache.Get("kkkk"))
	//cache.Set("kkkk", "aaaaa")
	//fmt.Println(cache.Get("kkkk"))
	//fmt.Println(cache.Get("hello"))
	//fmt.Println(cache.Get("cccccc"))
	//var duration time.Duration
	//fmt.Println(time.Second + time.Microsecond + 5)
	//fmt.Println()
	//b := true
	//f := false
	//
	//pointer := unsafe.Pointer(&b)
	//fmt.Printf("\n%t\n",pointer)
	//swapped := atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&f)), unsafe.Pointer(&f), unsafe.Pointer(&b))
	//if swapped {
	//	fmt.Println("hahahha")
	//}
	//fmt.Println(f)
	//
	//a := 55
	//atomic.
	ch := make([]chan interface{}, 4)
	for i := 0; i < 4; i++ {
		ch[i] = make(chan interface{}, 55)
	}
	fmt.Println(ch)
}
