package algorithm

import (
	"container/heap"
)

type Comparable interface {
	CompareTo(other Comparable) int
}

type PriorityQueue struct {
	base *priorityQueue
}

func NewPriorityQueue() *PriorityQueue {
	priorityQueue := PriorityQueue{base:new(priorityQueue)}
	heap.Init(priorityQueue.base)
	return &priorityQueue
}


type priorityQueue []Comparable

func (pq *priorityQueue) Less(i, j int) bool {
	//return true
	//fmt.Println(pq.Len())
	return (*pq)[i].CompareTo((*pq)[j]) < 0
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Comparable))
}

func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	if n <= 0 {
		return nil
	}
	x := (*pq)[n - 1]
	*pq = (*pq)[0 : n-1]
	return x
}


func (pq *priorityQueue) Swap(i, j int)  {
	(*pq)[i],(*pq)[j] = (*pq)[j],(*pq)[i]
}

func (pq priorityQueue) Len() int {
	return len(pq)
}


func (pq *PriorityQueue) Push(x Comparable) {
	heap.Push(pq.base, x)
}

func (pq *PriorityQueue) Pop() Comparable {
	return heap.Pop(pq.base).(Comparable)
}

func (pq PriorityQueue) Top() Comparable {

	return (*pq.base)[0].(Comparable)
}


func (pq PriorityQueue) Length() int {
	return pq.base.Len()
}

