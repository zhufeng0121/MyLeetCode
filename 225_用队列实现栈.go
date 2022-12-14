package main

type MyStack struct {
	majorQueue  []int
	helperQueue []int
}

func ConstructorMyStack() MyStack {
	return MyStack{
		majorQueue:  make([]int, 0),
		helperQueue: make([]int, 0),
	}

}

func (this *MyStack) Push(x int) {
	this.helperQueue = append(this.helperQueue, x)

	for len(this.majorQueue) > 0 {
		this.helperQueue = append(this.helperQueue, this.majorQueue[0])
		this.majorQueue = this.majorQueue[1:]
	}

	this.majorQueue, this.helperQueue = this.helperQueue, this.majorQueue

}

func (this *MyStack) Pop() int {
	v := this.majorQueue[0]
	this.majorQueue = this.majorQueue[1:]
	return v

}

func (this *MyStack) Top() int {
	return this.majorQueue[0]

}

func (this *MyStack) Empty() bool {
	return len(this.majorQueue) == 0

}
