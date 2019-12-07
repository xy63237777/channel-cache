package gocache

import (
	alg "four-seasons/algorithm"
	"time"
)




func (node needClearNode) CompareTo(other alg.Comparable) int {
	return int(node.expiration - other.(needClearNode).expiration)
}

func (node needClearNode) IsExpired() bool {
	if node.expiration <= int64(NoExpiration) {
		return false
	} else if node.expiration > time.Now().UnixNano() {
		return false
	}
	return true
}

func newLiquidator() *Liquidator {
	return &Liquidator{
		elements: alg.NewPriorityQueue(),
	}
}

func (l *Liquidator) clearNode(number int) (total int) {
	total = 0
	if l.elements.Length() == 0 {
		return
	}
	for i := 0; i < number; i++ {
		top := l.elements.Top().(*needClearNode)
		if top.IsExpired() {
			top = l.elements.Pop().(*needClearNode)
			_ = top.masterCache.DeleteAsync(top.key)
		} else {
			break
		}
		total++
	}
	return
}

func (l *Liquidator) clearFunc() {
	num := 0
	if l.elements.Length() < 50 {
		num = l.elements.Length()
	} else if l.elements.Length() > 1000 {
		num = 100
	} else {
		num = l.elements.Length() >> 1
	}
	l.clearNode(num)
}