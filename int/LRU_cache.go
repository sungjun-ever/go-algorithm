package main

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func NewLRUCache(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]

	if !ok {
		return -1
	}

	this.moveToHead(node)
	return node.val

}

func (this *LRUCache) Put(key int, value int) {
	node, exists := this.cache[key]
	if exists {
		node.val = value
		this.moveToHead(node)
		return
	}

	// 새로운 노드 생성 및 추가
	newNode := &Node{key: key, val: value}
	this.cache[key] = newNode
	this.addToHead(newNode)

	// 용량 초과 시 가장 오래된 노드(Tail 직전 노드) 제거
	if len(this.cache) > this.capacity {
		removed := this.popTail()
		delete(this.cache, removed.key)
	}
}

// 노드를 리스트에서 끊어내는 연산
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 항상 Dummy Head 바로 뒤에 노드를 삽입 (가장 최근 사용 표시)
func (this *LRUCache) addToHead(node *Node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

// 기존 노드의 위치를 맨 앞으로 갱신
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

// 가장 오랫동안 사용되지 않은 Dummy Tail 직전 노드를 삭제하고 반환
func (this *LRUCache) popTail() *Node {
	res := this.tail.prev
	this.removeNode(res)
	return res
}
