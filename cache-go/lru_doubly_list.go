package gocache

func newLRUChain() *lruNodeChain {
	head := &lruNode{}
	tail := &lruNode{}
	head.next = tail
	tail.prev = head
	return &lruNodeChain{
		head:head,
		tail:tail,
	}
}

func (ln *lruNode) LeaveForList() {
	ln.prev.next = ln.next
	ln.next.prev = ln.prev
	ln.next = nil
	ln.prev = nil
}

func (lc *lruNodeChain) IsEmpty() bool {
	return lc.head.next == lc.tail
}

func (lc *lruNodeChain) GetHead() *lruNode {
	if lc.head.next == lc.tail  {
		return nil
	}
	return lc.head.next
}

func (lc *lruNodeChain) GetTail() *lruNode {
	if lc.head.next == lc.tail  {
		return nil
	}
	return lc.tail.prev
}

func (lc *lruNodeChain) AddFirst(node *lruNode)  {
	node.prev = lc.head
	node.next = lc.head.next
	lc.head.next.prev = node
	lc.head.next = node
}

func (lc *lruNodeChain) RemoveHead() (*lruNode, bool) {
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

func (lc *lruNodeChain) AddTail(node *lruNode) {
	node.next = lc.tail
	node.prev = lc.tail.prev
	lc.tail.prev.next = node
	lc.tail.prev = node
}

func (lc *lruNodeChain) RemoveTail() (*lruNode, bool) {
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
