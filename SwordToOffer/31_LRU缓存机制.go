package SwordToOffer

type Node struct {
	prev  *Node
	post  *Node
	key   int
	value int
}

type DoubleList struct {
	size int
	head *Node
	tail *Node
}

func NewDoubleList() *DoubleList {
	list := &DoubleList{
		size: 0,
		head: &Node{},
		tail: &Node{},
	}

	list.head.post = list.tail
	list.tail.prev = list.head
	return list
}

// 节点添加到末尾
func (l *DoubleList) addLast(node *Node) {
	node.prev = l.tail.prev
	node.post = l.tail
	l.tail.prev.post = node
	l.tail.prev = node

	l.size++

}

func (l *DoubleList) remove(node *Node) {
	node.prev.post = node.post
	node.post.prev = node.prev

	l.size--
}

// 返回删除的头结点的
func (l *DoubleList) removeFirst() *Node {
	if l.head.post == l.tail {
		return nil
	}
	first := l.head.post
	l.remove(first)
	return first

}

func (l *DoubleList) Size() int {
	return l.size
}

type LRUCache struct {
	cap  int
	hash map[int]*Node
	list *DoubleList
}

func ConstructorV(capacity int) LRUCache {
	return LRUCache{
		cap:  capacity,
		hash: map[int]*Node{},
		list: NewDoubleList(),
	}

}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hash[key]
	if !ok {
		return -1
	}
	this.list.remove(node)
	this.list.addLast(node)
	return node.value

}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.hash[key]
	if ok {
		// 更新
		this.list.remove(node)
		node.value = value
		this.list.addLast(node)
		return
	}

	if this.list.Size() == this.cap {
		head := this.list.removeFirst()
		delete(this.hash, head.key)
	}

	newNode := &Node{
		key:   key,
		value: value,
	}
	this.list.addLast(newNode)
	this.hash[key] = newNode

}
