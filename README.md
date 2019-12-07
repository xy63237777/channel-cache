# channel-cache (Golang)

[toc]



![img](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1575743108362&di=5d0f84a64ea0b46d46fab8046b06664b&imgtype=0&src=http%3A%2F%2Fpic2.zhimg.com%2Fv2-e23145800bbd3d684aef85ad51145eee_1200x500.jpg)

## 引言

- ​      这是一款高性能的Go语言的缓存框架.笔者写的并不完美.虽然是线程安全并且使用了无锁.基于管道命令.支持同步异步获取.使用了LRU和LFU两种淘汰策略并且可以设置缓存时间.

- ​      然后这里我会进行一个快速开始.并且教你们如何使用.如果你想学习这个缓存项目的源代码并且提高功力.或者你可以帮帮笔者继续提高一下性能.那么我可以为你提供博客的教学地址.并且我会在博客的教学地址里写上我当时的完整思路还有问题的所在.如果解决这个问题的.

​			博客地址: 

-  **如果你是想学习Go语言没有好的项目并且想提高功力,那么这个项目一定是你的首选**

## 快速开始



你只需要下面这条命令就可以把这个正式版的项目copy下来

```shell
go get -v github.com/xy63237777/channel-cache
```

如果你想学习源码或者能为这个项目提交你的思路和代码 你可以切换至dev分支

![image-20191207234631095](/home/thesevensky/.config/Typora/typora-user-images/image-20191207234631095.png)



然后在你的代码里获取到一个Cache实例

```
package main

import (
	"github.com/xy63237777/channel-cache/cache-goV1"
	
)

func main() {
	cache := gocache.NewSignalCacheForDefault()
}
```

