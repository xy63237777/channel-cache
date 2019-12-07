package algorithm

type DoublyListNode struct {
	prev *DoublyListNode
	next *DoublyListNode
	Data interface{}
}


type DoublyLinkedList struct {
	size int
	head *DoublyListNode
	tail *DoublyListNode
}

func New() *DoublyLinkedList {
	head := &DoublyListNode{}
	tail := &DoublyListNode{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{
		head:head,
		tail:tail,
		size:0,
	}
}

func (dln *DoublyListNode) LeaveChain() {
	dln.next.prev = dln.prev
	dln.prev.next = dln.next
	dln.next = nil
	dln.prev = nil
}

func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

func (dll *DoublyLinkedList) GetHead() interface{} {
	if dll.Size() == 0 {
		return nil
	}
	return dll.head.next.Data
}

func (dll *DoublyLinkedList) GetHeadForNode() *DoublyListNode {
	if dll.Size() == 0 {
		return nil
	}
	return dll.head.next
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.head.next == dll.tail
}

func (dll *DoublyLinkedList) GetTailForNode() *DoublyListNode {
	if dll.Size() == 0 {
		return nil
	}
	return dll.tail.next
}

func (dll *DoublyLinkedList) AddFirst(data interface{})  {
	node := &DoublyListNode{Data:data, next:dll.head.next, prev:dll.head}
	dll.head.next.prev = node
	dll.head.next = node
	dll.size++
}

func (dll *DoublyLinkedList) RemoveHead() (interface{}, bool) {
	if dll.Size() == 0 {
		return nil, false
	}
	node := dll.head.next
	node.next.prev = dll.head
	dll.head.next = node.next
	node.next = nil
	node.prev = nil
	return node.Data, true
}

func (dll *DoublyLinkedList) RemoveHeadForNode() (*DoublyListNode, bool) {
	if dll.Size() == 0 {
		return nil, false
	}
	node := dll.head.next
	node.next.prev = dll.head
	dll.head.next = node.next
	node.next = nil
	node.prev = nil
	return node, true
}

func (dll *DoublyLinkedList) AddTail(data interface{}) {
	node := &DoublyListNode{Data:data, next:dll.tail, prev:dll.tail.prev}
	dll.tail.prev.next = node
	dll.tail.prev = node
	dll.size++
}

func (dll *DoublyLinkedList) RemoveTail() (interface{}, bool) {
	if dll.Size() == 0 {
		return nil, false
	}
	node := dll.tail.prev
	node.prev.next = dll.tail
	dll.tail.prev = node.prev
	node.prev = nil
	node.next = nil
	return node.Data, true
}

func (dll *DoublyLinkedList) RemoveTailForNode() (*DoublyListNode, bool) {
	if dll.Size() == 0 {
		return nil, false
	}
	node := dll.tail.prev
	node.prev.next = dll.tail
	dll.tail.prev = node.prev
	node.prev = nil
	node.next = nil
	return node, true
}

