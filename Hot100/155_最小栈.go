package Hot100

type MinStack struct {
	arr    []int
	helper []int // 辅助栈，用来记录最小值
}

func ConstructorI() MinStack {
	return MinStack{
		arr:    make([]int, 0),
		helper: make([]int, 0),
	}

}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	if len(this.helper) == 0 {
		this.helper = append(this.helper, val)
	} else {
		top := this.helper[len(this.helper)-1]
		if val < top {
			this.helper = append(this.helper, val)
		} else {
			this.helper = append(this.helper, top)
		}
	}

}

func (this *MinStack) Pop() {
	if len(this.arr) > 0 {
		this.arr = this.arr[:len(this.arr)-1]
		this.helper = this.helper[:len(this.helper)-1]
	}

}

func (this *MinStack) Top() int {
	var top int
	if len(this.arr) > 0 {
		top = this.arr[len(this.arr)-1]
	}
	return top

}

func (this *MinStack) GetMin() int {
	if len(this.arr) > 0 {
		return this.helper[len(this.helper)-1]
	}
	return 0

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
