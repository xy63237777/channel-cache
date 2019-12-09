package gocache

import (
	"four-seasons/log"
	"testing"
)

func TestCache_IsClose(t *testing.T) {
	ch := make(chan interface{},5)
	ch <- "hello"
	ch <- "world"
	ch <- "common"
	close(ch)
	for lb := range ch {
		log.Warn("hello %s", lb)
	}
}