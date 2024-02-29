package main

// LinkedListNode 链表节点结构体
type LinkedListNode struct {
	key  int
	val  int
	pre  *LinkedListNode
	next *LinkedListNode
}

// LinkedList 链表结构体
type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
}

// LRUCache LRU缓存结构体
type LRUCache struct {
	kv       map[int]*LinkedListNode
	list     *LinkedList
	capacity int
}

// NewLRUCache 初始化LRUCache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		kv:       make(map[int]*LinkedListNode),
		list:     new(LinkedList),
		capacity: capacity,
	}
}

// AddHead 在链表头部添加节点
func (l *LinkedList) AddHead(node *LinkedListNode) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	node.next = l.head
	l.head.pre = node
	l.head = node
}

// DelNode 删除节点
func (l *LinkedList) DelNode(node *LinkedListNode) {
	if node.pre != nil {
		node.pre.next = node.next
	}
	if node.next != nil {
		node.next.pre = node.pre
	}
	if node == l.head {
		l.head = node.next
	}
	if node == l.tail {
		l.tail = node.pre
	}
}

// Get 获取LRU缓存中的值
func (l *LRUCache) Get(key int) int {
	if node, ok := l.kv[key]; ok {
		if node != l.list.head {
			l.list.DelNode(node)
			l.list.AddHead(node)
		}
		return node.val
	}
	return -1
}

// Put 向LRU缓存中添加值
func (l *LRUCache) Put(key, value int) {
	if node, ok := l.kv[key]; ok {
		l.list.DelNode(node)
	} else if len(l.kv) >= l.capacity {
		delete(l.kv, l.list.tail.key)
		l.list.DelNode(l.list.tail)
	}
	node := &LinkedListNode{key: key, val: value}
	l.list.AddHead(node)
	l.kv[key] = node
}
