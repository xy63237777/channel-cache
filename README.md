# channel-cache (Golang)

[toc]



![img](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1575743108362&di=5d0f84a64ea0b46d46fab8046b06664b&imgtype=0&src=http%3A%2F%2Fpic2.zhimg.com%2Fv2-e23145800bbd3d684aef85ad51145eee_1200x500.jpg)

## 引言

- ​      这是一款高性能的Go语言的缓存框架.笔者写的并不完美.虽然是线程安全并且使用了无锁.基于管道命令.支持同步异步获取.使用了LRU和LFU两种淘汰策略并且可以设置缓存时间.

- ​      然后这里我会进行一个快速开始.并且教你们如何使用.如果你想学习这个缓存项目的源代码并且提高功力.或者你可以帮帮笔者继续提高一下性能.那么我可以为你提供博客的教学地址.并且我会在博客的教学地址里写上我当时的完整思路还有问题的所在.如果解决这个问题的.

​			博客地址: 这里还正在写 别着急~

-  **如果你是想学习Go语言没有好的项目并且想提高功力,那么这个项目一定是你的首选**

## 快速开始



你只需要下面这条命令就可以把这个正式版的项目copy下来

```shell
go get -v github.com/xy63237777/channel-cache
```

如果你想学习源码或者能为这个项目提交你的思路和代码 你可以切换至dev分支

dev : https://github.com/xy63237777/channel-cache/tree/channel-dev



然后在你的代码里获取到一个Cache实例

```go
package main

import (
	"github.com/xy63237777/channel-cache/cache-goV1"
	
)

func main() {
	cache := gocache.NewSignalCacheForDefault()
}

```

当然你也可以指定一些策略

第一个参数是LRU(最近最久未使用)缓存 默认是LFU(最近最少使用算法)  然后第二个参数是这个缓存的大小.

默认是2048

这是一个简单开始的例子

```go
cache := gocache.NewSignalCache(gocache.CacheForLRU, gocache.DefaultCapacity)
```

然后对于返回会返回值本身和bool类型. bool类型表示有无这个对象 通常会返回你想要的值+true 或者 nil + false

```go
func main() {
	cache := gocache.NewSignalCache(gocache.CacheForLRU, gocache.DefaultCapacity)
	cache.Set("hello","world")
	fmt.Println(cache.Get("hello"))
}
```

当然你在Set的时候为这个对象写一定的过期时间. 如果不指定则默认不过期

下面代码如果3个小时后就会过期

```go
func main() {
	cache := gocache.NewSignalCache(gocache.CacheForLRU, gocache.DefaultCapacity)
	cache.SetForExpiration("hello","world", time.Hour * 3)
	fmt.Println(cache.Get("hello"))
}
```

下面代码将异步返回你的结果

返回的先是一个管道 你通过这个管道再去读. 也就是说 你可以先去获管道,如果你不着急获取值的话. 你可以在客户端异步去完成这些工作. 等你真正需要这个变量的时候再通过管道去拿.

```go
func main() {
	cache := gocache.NewSignalCache(gocache.CacheForLRU, gocache.DefaultCapacity)
	cache.SetForExpiration("hello","world", time.Hour * 3)
	async := cache.GetAsync("hello")
	val :=  <- async
	fmt.Println(val)
}
```

当然你可以删除一个缓存实例

```go
func main() {
	cache := gocache.NewSignalCache(gocache.CacheForLRU, gocache.DefaultCapacity)
	cache.SetForExpiration("hello","world", time.Hour * 3)
	async := cache.GetAsync("hello")
	val :=  <- async
	fmt.Println(val)
	cache.Delete("hello")
}
```

## 较高级操作



你可以创建一个缓存管理者,这个管理者是单例的.所以放心使用

如果你想单独做一个缓存服务器的话,你可以制定一个go的服务端和客户端来做一些负载均衡.,那么就应该使用manager 如果你只需要一个简单的缓存NewSignalCache应该是你的首选

```go
func main() {
	manager := gocache.NewCacheManager()
	manager.CreateCacheForDefault("cache1")
}
```

然后通过Get这个缓存就和以前使用一样了

```go
func main() {
	manager := gocache.NewCacheManager()
	manager.CreateCacheForDefault("cache1")
	cache := manager.GetCache("cache1")
	cache.Set("hello", "world")
	cache.Get("hello")
}
```

你可以暂停或者继续一个缓存

如果当你暂停以后你的所有操作都不会接受. 并且会打印warning

```go
func main() {
	manager := gocache.NewCacheManager()
	manager.CreateCacheForDefault("cache1")
	cache := manager.GetCache("cache1")
	err := manager.CacheStop("cache1")
	if err != nil {
		panic(err)
	}
	cache.Set("hello", "world")
	fmt.Println(cache.Get("hello"))
	err = manager.CacheRun("cache1")
	if err != nil {
		panic(err)
	}
	cache.Set("hello", "world")
	fmt.Println(cache.Get("hello"))
}
```

以下是笔者电脑上的打印信息

```
[GOPM] 12-08 00:21:23 [ WARN] Warning...  Using a stopped cache
[GOPM] 12-08 00:21:23 [ WARN] Warning...  Using a stopped cache
<nil> false
world true

```

当然你也可以像管理者里加入一个缓存实例或者删除一个实例