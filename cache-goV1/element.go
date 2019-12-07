package gocache

import (
	"time"
)

//func (item Item) CompareTo(other alg.Comparable) int {
//	return int(item.expiration - other.(Item).expiration)
//}

func (item item) IsExpired() bool {
	if item.expiration <= int64(NoExpiration) {
		return false
	} else if item.expiration > time.Now().UnixNano() {
		return false
	}
	return true
}
func newItem(obj interface{}, expiration time.Duration) *item {
	exp := time.Now().UnixNano() + expiration.Nanoseconds()
	if expiration == NoExpiration {
		exp = int64(NoExpiration)
	}
	return &item{
		obj:        obj,
		expiration: exp,
	}
}

func newClearNode(ca *Cache, key string, expiration time.Duration) *needClearNode {
	exp := time.Now().UnixNano() + expiration.Nanoseconds()
	return &needClearNode{
		masterCache: ca,
		key:         key,
		expiration:  exp,
	}
}

