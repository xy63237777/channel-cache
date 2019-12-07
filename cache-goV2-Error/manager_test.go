package gocache

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewCacheManger(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		group.Add(1)
		go func() {

			fmt.Printf("%p\n",NewCacheManager())
			group.Done()
		}()
	}
	group.Wait()
}


