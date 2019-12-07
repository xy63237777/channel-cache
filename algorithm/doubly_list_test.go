package algorithm

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	list := New()
	str := "hello"
	list.AddFirst(str)
	list.AddTail("abc")
	list.AddFirst("kkkk")
	show(list)
	list.RemoveTail()
	show(list)
	list.RemoveHead()
	show(list)
}

func show(list *DoublyLinkedList)  {
	for temp := list.head.next; temp != nil; temp = temp.next {
		//fmt.Println(temp.data)
	}
}