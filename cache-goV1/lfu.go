package gocache

func (lc *lfuCache) Size() int {
	return lc.size
}

func newLFUCache(capacity int, master *Cache) *lfuCache {
	lfuCache := &lfuCache{
		capacity: capacity,
		size:     0,
		elements: make(map[string]*lfuNode, (capacity<<1|capacity)+1),
		manager: lfuChainManager{
			firstLinkedList: newLFUChain(0),
			lastLinkedList:  newLFUChain(0),
		},
		master:master,
	}
	lfuCache.manager.firstLinkedList.next = lfuCache.manager.lastLinkedList
	lfuCache.manager.lastLinkedList.prev = lfuCache.manager.firstLinkedList
	return lfuCache
}

func (lc *lfuCache) Capacity() int {
	return lc.capacity
}

func (lc *lfuCache) Get(key string) interface{} {
	node, ok := lc.elements[key]
	if !ok {
		return nil
	}
	freqInc(node)
	return node.data
}

func (lc *lfuCache) Delete(key string)  {
	if lc.capacity <= 0 {
		return
	}
	node,ok := lc.elements[key]
	if !ok {
		return
	}
	lc.deleteNode(node)
	delete(lc.elements,key)
	lc.size--

}

func (lc *lfuCache) deleteNode(node *lfuNode)  {
	chain := node.parent
	node.LeaveForList()
	if chain.IsEmpty() {
		chain.LeaveForChain()
	}
}


func (lc *lfuCache) Put(key string, value interface{})  {
	if lc.capacity <= 0 {
		return
	}
	node , ok := lc.elements[key]
	if ok {
		node.data = value
		freqInc(node)
	} else {
		if lc.Size() >= lc.capacity {

			listNodeChain := lc.manager.firstLinkedList.next
			delNode := listNodeChain.GetTail()
			lc.deleteNode(delNode)
			delete(lc.elements, delNode.key)
			//delNode, _ := listNodeChain.GetTail()
			//if listNodeChain.IsEmpty() {
			//	listNodeChain.LeaveForChain()
			//}
			lc.size--
		}
		lfuNode := &lfuNode{
			key:key,
			data:value,
			freq:1,
		}
		lc.elements[key] = lfuNode
		if lc.manager.firstLinkedList.next.freq != 1 {
			tempList := newLFUChain(1)
			tempList.AddFirst(lfuNode)
			insertChain(lc.manager.firstLinkedList, tempList)
		} else {
			lc.manager.firstLinkedList.next.AddFirst(lfuNode)
		}
		lc.size++
	}
}

func freqInc(node *lfuNode)  {
	parentList := node.parent
	node.LeaveForList()
	nextList := parentList.next
	tempParent := parentList.prev
	if parentList.IsEmpty() {
		parentList.LeaveForChain()
		parentList = tempParent
	}
	node.freq++
	if nextList.freq != node.freq {
		newChain := newLFUChain(node.freq)
		insertChain(parentList, newChain)
		newChain.AddFirst(node)
		node.parent = newChain
	} else {
		nextList.AddFirst(node)
	}
}

func insertChain(preChain, newChain *lfuNodeChain)  {
	newChain.next = preChain.next
	preChain.next.prev = newChain
	newChain.prev = preChain
	preChain.next = newChain
}