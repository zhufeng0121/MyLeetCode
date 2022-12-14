package SwordToOffer

type MinStack struct {
	major  []int
	helper []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		major:  []int{},
		helper: []int{},
	}

}

func (this *MinStack) Push(x int) {
	this.major = append(this.major, x)

	if len(this.helper) > 0 && x > this.helper[len(this.helper)-1] {
		this.helper = append(this.helper, this.helper[len(this.helper)-1])
	} else {
		this.helper = append(this.helper, x)
	}

}

func (this *MinStack) Pop() {
	if len(this.major) > 0 {
		this.major = this.major[:len(this.major)-1]
		this.helper = this.helper[:len(this.helper)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.major) > 0 {
		return this.major[len(this.major)-1]
	}
	return -1

}

func (this *MinStack) Min() int {
	return this.helper[len(this.helper)-1]

}