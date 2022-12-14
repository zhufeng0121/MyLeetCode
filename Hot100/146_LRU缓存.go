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
	Head *Node
	Tail *Node
}

// Push 向双链表的尾部添加新节点
func (list *DoubleList) Push(node *Node) {
	// 如果双链表为空
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		// 更改尾节点
		list.Tail = node
		list.Tail.post = node
		node.prev = list.Tail

	}

}

// Pop 从双链表的头部弹出节点
func (list *DoubleList) Pop() {

}

// Splice 从双链表中抽离节点
func (list *DoubleList) Splice(node *Node) {

}

// LRUCache 采用 双链表和 map结构来实现
type LRUCache struct {
	Len  int           // LRU 目前的长度
	Cap  int           // LRU 的总容量
	Hash map[int]*Node // 根据值返回相应的节点
	List *DoubleList   // LRU的链表
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Len:  0,
		Cap:  capacity,
		Hash: map[int]*Node{},
		List: &DoubleList{
			Head: nil,
			Tail: nil,
		},
	}
}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
