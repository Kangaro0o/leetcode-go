package design

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

type LRUCache struct {
	cache                map[int]*LinkedNode
	size, capacity       int
	dummyHead, dummyTail *LinkedNode
}

func Constructor1(capacity int) LRUCache {
	head := &LinkedNode{key: -1, value: -1}
	tail := &LinkedNode{key: -1, value: -1}
	head.next = tail
	tail.prev = head
	return LRUCache{
		size:      0,
		capacity:  capacity,
		cache:     make(map[int]*LinkedNode),
		dummyHead: head,
		dummyTail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	// 移动到队头
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok { // 存在
		node.value = value
		this.moveToHead(node)
	} else {
		n := &LinkedNode{
			key:   key,
			value: value,
		}
		this.addHead(n)
		this.cache[key] = n
		this.size++
		if this.size > this.capacity {
			curr := this.dummyTail.prev
			delete(this.cache, curr.key)
			this.delTail()
			this.size--
		}
	}
}

func (this *LRUCache) addHead(node *LinkedNode) {
	// 添加到队列头
	first := this.dummyHead.next
	this.dummyHead.next = node
	node.prev = this.dummyHead
	node.next = first
	first.prev = node
}

func (this *LRUCache) delTail() {
	curr := this.dummyTail.prev
	prev := curr.prev
	prev.next = this.dummyTail
	this.dummyTail.prev = prev
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
	// 删除 node
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev

	// 添加到队列头
	this.addHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
