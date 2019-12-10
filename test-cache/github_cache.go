package main

import (
	"fmt"
	ca "four-seasons/cache-goV1"
	"github.com/patrickmn/go-cache"
	"strconv"
	"sync"
	"time"
)
var N = 300000
var M = 8
func main() {
	manager := ca.NewCacheManager()
	fmt.Println(manager)
	testForGit()
	//testForMy2()
	testForMy()
	testForMyG()
	testForMyMore()
	//testForMyLRU()
	testForGitPlus()
	//testForMyPlus2()
	//testForMyPlus()
	testForMyLRUPlus()

}

func testForGit()  {
	cache := cache.New(30*time.Second, 10*time.Second)
	start := time.Now().UnixNano()
	for i := 0; i < N; i++{
		k := i
		cache.Set(strconv.Itoa(k), strconv.Itoa(i), time.Hour)
		cache.Get(strconv.Itoa(k))
	}
	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
}

func testForMy()  {
	cache := ca.NewSignalCache(ca.CacheForLFU, 4096)
	fmt.Println(cache)
	start := time.Now().UnixNano()
	for i := 0; i < N; i++{
		k := i
		cache.Set(strconv.Itoa(k), strconv.Itoa(i))
		cache.GetAsync(strconv.Itoa(k))
	}
	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
}

func testForMyG()  {
	cache := ca.NewSignalCache(ca.CacheForLFU, 4096)
	fmt.Println(cache)
	start := time.Now().UnixNano()
	for i := 0; i < N; i++{
		k := i
		cache.SetForExpiration(strconv.Itoa(k), strconv.Itoa(i),time.Nanosecond)
		cache.GetAsync(strconv.Itoa(k))
	}
	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
	time.Sleep(time.Second * 5)
}

func testForMyMore()  {
	cache := ca.NewSignalCache(ca.CacheForLFU, 4096)
	fmt.Println(cache)
	manager := ca.NewCacheManager()
	cacheArr := []*ca.Cache{}
	for i := 0; i < 4; i++ {
		manager.CreateCacheForDefault(strconv.Itoa(i))
		cacheArr = append(cacheArr, manager.GetCache(strconv.Itoa(i)))
	}
	start := time.Now().UnixNano()
	for i := 0; i < N; i++{
		k := i
		cacheArr[i % 4].Set(strconv.Itoa(k), strconv.Itoa(i))
		cacheArr[i % 4].Get(strconv.Itoa(k))
	}
	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
}

//func testForMy2()  {
//	cache := ca.NewSignalCache(ca.CacheForLRU, 4096)
//	start := time.Now().UnixNano()
//	for i := 0; i < N; i++{
//		k := i % 4096
//		cache.Set(strconv.Itoa(k), strconv.Itoa(i))
//		cache.Get(strconv.Itoa(k))
//	}
//	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
//}

func testForMyLRU()  {
	cache := ca.NewSignalCache(ca.CacheForLRU, ca.DefaultCapacity)
	group := sync.WaitGroup{}
	start := time.Now().UnixNano()
	for i := 0; i < M; i++ {
		group.Add(1)
		go func() {

			for i := 0; i < N; i++{
				k := i
				cache.Set(strconv.Itoa(k), strconv.Itoa(i))
				cache.Get(strconv.Itoa(k))
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
}

func testForGitPlus()  {
	cache := cache.New(30*time.Second, 10*time.Second)
	group := sync.WaitGroup{}
	start := time.Now().UnixNano()
	for i := 0; i < M; i++ {

		group.Add(1)
		go func() {

			for j := 0; j < N; j++{
				k := j
				cache.Set(strconv.Itoa(k), strconv.Itoa(i),time.Hour)
				cache.Get(strconv.Itoa(k))
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
}

func testForMyPlus()  {
	//cache := ca.NewSignalCacheForDefault()
	manager := ca.NewCacheManager()
	fmt.Println(manager)
	for i := 0; i < M; i++ {
		manager.CreateCacheForDefault(strconv.Itoa(i))
	}

	group := sync.WaitGroup{}
	start := time.Now().UnixNano()
	for i := 0; i < M; i++ {
		group.Add(1)
		go func() {

			for j := 0; j < N; j++{
				k := j
				cache := manager.GetCache(strconv.Itoa(j % M))
				cache.Set(strconv.Itoa(k), strconv.Itoa(i))
				cache.Get(strconv.Itoa(k))
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
}
//
func testForMyPlus2()  {
	cache := ca.NewSignalCacheForDefault()
	group := sync.WaitGroup{}
	start := time.Now().UnixNano()
	for i := 0; i < M; i++ {
		group.Add(1)
		go func() {

			for j := 0; j < N; j++{
				k := j
				cache.Set(strconv.Itoa(k), strconv.Itoa(i))
				cache.Get(strconv.Itoa(k))
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
}
//
func testForMyLRUPlus()  {
	cache := ca.NewSignalCache(ca.CacheForLFU, 22222)
	group := sync.WaitGroup{}
	start := time.Now().UnixNano()
	for i := 0; i < M; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < N/10; j++{
				k := j
				cache.Set(strconv.Itoa(k), strconv.Itoa(i))
				cache.GetAsync(strconv.Itoa(k))
			}
			fmt.Println("hello123")
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println((time.Now().UnixNano() - start)/int64(time.Millisecond))
}