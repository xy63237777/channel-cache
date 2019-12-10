package gocache

import (
	alg "four-seasons/algorithm"
	"four-seasons/log"
	"sync"
	"time"
)

func (node needClearNode) CompareTo(other alg.Comparable) int {
	return int(node.expiration - other.(*needClearNode).expiration)
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
		mu: sync.RWMutex{},
	}
}

func (l *Liquidator) top() alg.Comparable {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.elements.Top()
}

func (l *Liquidator) push(obj alg.Comparable) {
	l.mu.Lock()
	l.elements.Push(obj)
	l.mu.Unlock()
}

func (l *Liquidator) pop() alg.Comparable {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.elements.Pop()
}

func (l *Liquidator) clearNode(number int) (total int) {
	total = 0
	n := l.elements.Length()
	if n == 0 {
		return
	}
	for i := 0; i < number; i++ {
		top := l.top().(*needClearNode)
		if top.IsExpired() {
			top = l.pop().(*needClearNode)
			top.masterCache.Delete(top.key)
		} else {
			break
		}
		total++
	}
	log.Info("clear Expired Node size is %d Node total Size is %d ",total, n)
	return
}

func (l *Liquidator) clearFunc() {
	num := 0
	if l.elements.Length() < 25 {
		num = l.elements.Length()
	} else if l.elements.Length() > 500 {
		num = 100
	} else {
		num = l.elements.Length() >> 2
	}
	l.clearNode(num)
}