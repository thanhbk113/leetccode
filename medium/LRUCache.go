package main

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	Capacity int
	Map      map[int]*Node
	Head     *Node
	Tail     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Map:      make(map[int]*Node, capacity),
		Head:     &Node{},
		Tail:     &Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Map[key]; ok {
		removeNode(node)
		this.addToHead(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) addToHead(node *Node) {
	if this.Head.Next == nil {
		this.Head.Next = node
		node.Prev = this.Head
		this.Tail.Prev = node
		node.Next = this.Tail
	} else {
		nodeBeforeHeadChange := this.Head.Next
		node.Prev = this.Head
		this.Head.Next = node
		node.Next = nodeBeforeHeadChange
		nodeBeforeHeadChange.Prev = node
	}
}

func removeNode(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Next = nil
	node.Prev = nil
	node = nil
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Map[key]; ok {
		node.Value = value
		removeNode(node)
		this.addToHead(node)
	} else {
		if len(this.Map) >= this.Capacity {
			delete(this.Map, this.Tail.Prev.Key)
			removeNode(this.Tail.Prev)
		}
		newNode := &Node{Key: key, Value: value}
		this.Map[key] = newNode
		this.addToHead(newNode)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	a := Constructor(2)

	a.Put(1, 1)
	a.Put(2, 2)
	a.Put(3, 3)
}
