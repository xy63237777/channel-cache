package gocache

func (lc *LRUCache) Size() int {
	return lc.size
}

func NewLRUCache(capacity int, master *Cache) *LRUCache {
	lruCache := &LRUCache{
		capacity: capacity,
		size:     0,
		elements: make(map[string]*lruNode, (capacity<<1|capacity)+1),
		manager: newLRUChain(),
		master:master,
	}
	return lruCache
}


func (lc *LRUCache) Get(key string) interface{} {
	node, ok := lc.elements[key]
	if !ok {
		return nil
	}
	lruUp(node, lc.manager)
	return node.data
}

func (lc *LRUCache) Delete(key string) interface{} {
	if lc.capacity <= 0 {
		return nil
	}
	node,ok := lc.elements[key]
	if !ok {
		return nil
	}
	lc.deleteNode(node)
	delete(lc.elements,key)
	lc.size--
	return node.data
}

func (lc *LRUCache) deleteNode(node *lruNode)  {
	node.LeaveForList()
}

func (lc *LRUCache) Put(key string, value interface{}) interface{} {
	if lc.capacity <= 0 {
		return nil
	}
	var oldVal interface{} = nil
	node , ok := lc.elements[key]
	if ok {
		oldVal = node.data
		node.data = value
		lruUp(node, lc.manager)
	} else {
		if lc.Size() >= lc.capacity {
			//先删除超时的不成功再删除LRU
			if lc.master.liquidator.clearNode(1) <= 0 {
				delNode, _ := lc.manager.RemoveHead()
				delete(lc.elements, delNode.key)
				lc.size--
			}

		}
		lruNode := &lruNode{
			key:key,
			data:value,
		}
		lc.elements[key] = lruNode
		lc.manager.AddTail(lruNode)
		lc.size++
	}
	return oldVal
}

func lruUp(node *lruNode, chain *lruNodeChain)  {
	node.LeaveForList()
	chain.AddTail(node)
}
