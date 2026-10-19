type ListNode struct{
	Key, Value int
	Pre, Next *ListNode
}

type LRUCache struct {
    kv map[int]*ListNode
	head, tail *ListNode
	capacity int
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
		kv: make(map[int]*ListNode),
		head: &ListNode{},
		tail: &ListNode{},
		capacity: capacity,
	}
	lru.head.Next = lru.tail
	lru.tail.Pre = lru.head
	return lru
}

func(this *LRUCache) remove(node *ListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func(this *LRUCache) setHead(node *ListNode) {
	node.Next = this.head.Next
	node.Pre = this.head
	this.head.Next.Pre = node
	this.head.Next = node
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.kv[key]; ok {
		node := this.kv[key]
		this.remove(node)
		this.setHead(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if _, ok := this.kv[key]; ok {
		node := this.kv[key]
		node.Value = value
		this.remove(node)
		this.setHead(node)
	} else {
		node := &ListNode{
			Key: key,
			Value: value,
		}
		if len(this.kv) == this.capacity {
			//delete
			del := this.tail.Pre
			this.remove(del)
			delete(this.kv, del.Key)
		} 
		this.setHead(node)
		this.kv[key] = node
	}
}
