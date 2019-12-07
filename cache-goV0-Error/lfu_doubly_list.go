package gocache

func (dln *lfuNode) LeaveForList()  {
	dln.prev.next = dln.next
	dln.next.prev = dln.prev
	dln.next = nil
	dln.prev = nil
	dln.parent = nil
}

func (lc *lfuNodeChain) LeaveForChain()  {
	lc.prev.next = lc.next
	lc.next.prev = lc.prev
	lc.prev = nil
	lc.next = nil
}

func (lc *lfuNodeChain) IsEmpty() bool {
	return lc.head.next == lc.tail
}

func newLFUChain(freq int) *lfuNodeChain {
	head := &lfuNode{}
	tail := &lfuNode{}
	head.next = tail
	tail.prev = head
	return &lfuNodeChain{
		head:head,
		tail:tail,
		freq:freq,
	}
}

func (lc *lfuNodeChain) GetHead() *lfuNode {
	if lc.head.next == lc.tail  {
		return nil
	}
	return lc.head.next
}

func (lc *lfuNodeChain) GetTail() *lfuNode {
	if lc.head.next == lc.tail  {
		return nil
	}
	return lc.tail.prev
}

func (lc *lfuNodeChain) AddFirst(node *lfuNode)  {
	node.prev = lc.head
	node.next = lc.head.next
	lc.head.next.prev = node
	lc.head.next = node
	node.parent = lc
}

func (lc *lfuNodeChain) RemoveHead() (*lfuNode, bool) {
	if lc.head.next == lc.tail  {
		return nil, false
	}
	node := lc.head.next
	node.next.prev = lc.head
	lc.head.next = node.next
	node.next = nil
	node.prev = nil
	return node, true
}

func (lc *lfuNodeChain) AddTail(node *lfuNode) {
	node.next = lc.tail
	node.prev = lc.tail.prev
	lc.tail.prev.next = node
	lc.tail.prev = node
	node.parent = lc
}

func (lc *lfuNodeChain) RemoveTail() (*lfuNode, bool) {
	if lc.head.next == lc.tail  {
		return nil, false
	}
	node := lc.tail.prev
	node.prev.next = lc.tail
	lc.tail.prev = node.prev
	node.prev = nil
	node.next = nil
	return node, true
}


