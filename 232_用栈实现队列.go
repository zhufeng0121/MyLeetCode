package main

type MyQueue struct {
	majorStack  []int
	helperStack []int
}

func ConstructorMyQueue() MyQueue {
	return MyQueue{
		majorStack:  make([]int, 0),
		helperStack: make([]int, 0),
	}

}

func (this *MyQueue) Push(x int) {
	this.majorStack = append(this.majorStack, x)

}

func (this *MyQueue) Pop() int {
	var top int
	if len(this.helperStack) > 0 {
		top = this.helperStack[len(this.helperStack)-1]
		this.helperStack = this.helperStack[:len(this.helperStack)-1]
		return top
	}
	for len(this.majorStack) > 0 {
		this.helperStack = append(this.helperStack, this.majorStack[len(this.majorStack)-1])
		this.majorStack = this.majorStack[:len(this.majorStack)-1]
	}

	top = this.helperStack[len(this.helperStack)-1]
	this.helperStack = this.helperStack[:len(this.helperStack)-1]
	return top

}

func (this *MyQueue) Peek() int {
	var top int
	if len(this.helperStack) > 0 {
		top = this.helperStack[len(this.helperStack)-1]
		return top
	}
	for len(this.majorStack) > 0 {
		this.helperStack = append(this.helperStack, this.majorStack[len(this.majorStack)-1])
		this.majorStack = this.majorStack[:len(this.majorStack)-1]
	}

	top = this.helperStack[len(this.helperStack)-1]
	return top

}

func (this *MyQueue) Empty() bool {
	if len(this.majorStack) == 0 && len(this.helperStack) == 0 {
		return true
	}
	return false

}
