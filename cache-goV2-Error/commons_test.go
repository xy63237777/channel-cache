package gocache

import (
	"context"
	"fmt"
	"testing"
)

func str(ctx context.Context)  {

	fmt.Println("hello")
}

func TestCache(t *testing.T) {
	//executor := &CacheExecutor{fn: (*func())(unsafe.Pointer(&str))}
	//executor.fn()
}