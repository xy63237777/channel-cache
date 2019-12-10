package gocache

func (lc *lruCache) Size() int {
	return lc.size
}

func newLRUCache(capacity int, master *Cache) *lruCache {
	lruCache := &lruCache{
		capacity: capacity,
		size:     0,
		elements: make(map[string]*lruNode, (capacity<<1|capacity)+1),
		manager: newLRUChain(),
		master:master,
	}
	return lruCache
}


func (lc *lruCache) Get(key string) interface{} {
	node, ok := lc.elements[key]
	if !ok {
		return nil
	}
	lruUp(node, lc.manager)
	return node.data
}

func (lc *lruCache) Delete(key string)  {
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

func (lc *lruCache) deleteNode(node *lruNode)  {
	node.LeaveForList()
}

func (lc *lruCache) Capacity() int {
	return lc.capacity
}

func (lc *lruCache) Put(key string, value interface{})  {
	if lc.capacity <= 0 {
		return
	}
	node , ok := lc.elements[key]
	if ok {
		node.data = value
		lruUp(node, lc.manager)
	} else {
		if lc.Size() >= lc.capacity {
			//先删除超时的不成功再删除LRU
				delNode, _ := lc.manager.RemoveHead()
				delete(lc.elements, delNode.key)
				lc.size--

		}
		lruNode := &lruNode{
			key:key,
			data:value,
		}
		lc.elements[key] = lruNode
		lc.manager.AddTail(lruNode)
		lc.size++
	}
}

func lruUp(node *lruNode, chain *lruNodeChain)  {
	node.LeaveForList()
	chain.AddTail(node)
}
