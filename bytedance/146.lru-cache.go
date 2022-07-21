package bytedance

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

type LRUCache struct {
	size, capacity       int
	cache                map[int]*LinkedNode
	dummyHead, dummyTail *LinkedNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkedNode{key: -1, value: -1}
	tail := &LinkedNode{key: -1, value: -1}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity:  capacity,
		size:      0,
		cache:     map[int]*LinkedNode{},
		dummyHead: head,
		dummyTail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		// 移动节点到链表头
		this.removeNode(node)
		this.moveToHead(node)
		//返回值
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		// 移动节点到链表头
		this.removeNode(node)
		this.moveToHead(node)
		node.value = value
		return
	}
	newNode := &LinkedNode{
		key:   key,
		value: value,
	}
	this.moveToHead(newNode)
	this.cache[key] = newNode
	this.size++
	if this.size > this.capacity {
		k := this.removeTail()
		delete(this.cache, k)
		this.size--
	}
}

func (this *LRUCache) removeTail() int {
	node := this.dummyTail.prev
	this.removeNode(node)
	return node.key
}

func (this *LRUCache) removeNode(node *LinkedNode) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
	first := this.dummyHead.next
	this.dummyHead.next = node
	node.prev = this.dummyHead
	node.next = first
	first.prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
