package algorithm

import (
	"fmt"
	"testing"
)

type myInt int

func (m myInt) CompareTo(other Comparable) int {
	return int(m - other.(myInt))
}

func TestNewPriorityQueue(t *testing.T) {
	var a myInt = 5
	var b myInt = 9
	var c myInt = 7
	queue := NewPriorityQueue()
	queue.Push(a)
	queue.Push(b)
	queue.Push(c)
	fmt.Println("top ", queue.Top())
	fmt.Println("top ", queue.Top())
	fmt.Println(queue.Pop())
	fmt.Println("top ", queue.Top())
	fmt.Println(queue.Pop())
	fmt.Println("top ", queue.Top())
	fmt.Println(queue.Pop())
}
