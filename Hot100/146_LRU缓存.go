package Hot100

// 双链表的节点
type Node struct {
	prev  *Node
	post  *Node
	key   int
	value int
}

// 双链表
type DoubleList struct {
	// 头尾节点 不过是虚节点
	Head *Node
	Tail *Node
	size int // 链表元素数
}

// 封装一些双链表的方法

func NewDoubleList() (list *DoubleList) {
	list = &DoubleList{
		Head: &Node{},
		Tail: &Node{},
		size: 0,
	}
	list.Head.post = list.Tail
	list.Tail.prev = list.Head
	return

}

// 双链表的尾部添加节点
func (list *DoubleList) addLast(node *Node) {
	node.prev = list.Tail.prev
	node.post = list.Tail
	list.Tail.prev.post = node
	list.Tail.prev = node

	list.size++

}

// 删除 node 节点 node节点一定存在
func (list *DoubleList) remove(node *Node) {
	node.prev.post = node.post
	node.post.prev = node.prev
	list.size--
}

// 删除链表中第一个节点，并返回该节点

func (list *DoubleList) removeFirst() (node *Node) {
	// 链表中没有元素
	if list.Head.post == list.Tail {
		return nil
	}

	first := list.Head.post
	list.remove(first)
	return first
}

func (list *DoubleList) Size() int {
	return list.size
}

// LRUCache 采用 双链表和 map结构来实现
type LRUCache struct {
	Cap  int           // LRU 的总容量
	Hash map[int]*Node // 根据值返回相应的节点
	List *DoubleList   // LRU的链表
}

func ConstructorII(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Hash: map[int]*Node{},
		List: NewDoubleList(),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Hash[key]
	if !ok {
		return -1
	}

	// 将该数据提升为最近使用的
	this.List.remove(node)
	this.List.addLast(node)

	return node.value

}

func (this *LRUCache) Put(key int, value int) {

	node, ok := this.Hash[key]
	// 更新node
	if ok {
		// 删除
		this.List.remove(node)

		// 更新node
		node.value = value
		this.List.addLast(node)
		return

	}

	if this.Cap == this.List.Size() {
		//删除头元素
		head := this.List.removeFirst()
		delete(this.Hash, head.key)
	}
	newNode := &Node{
		key:   key,
		value: value,
	}
	this.List.addLast(newNode)
	this.Hash[key] = newNode

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
