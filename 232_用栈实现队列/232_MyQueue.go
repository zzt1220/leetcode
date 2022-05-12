package _32_用栈实现队列

type MyQueue struct {
	stackIn  MyStack
	stackOut MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  MyStackConstructor(),
		stackOut: MyStackConstructor(),
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			x := this.stackIn.Pop()
			this.stackOut.Push(x)
		}
	}
	return this.stackOut.Pop()
}

func (this *MyQueue) Peek() int {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			x := this.stackIn.Pop()
			this.stackOut.Push(x)
		}
	}
	return this.stackOut.Top()
}

func (this *MyQueue) Empty() bool {
	return this.stackIn.Empty() && this.stackOut.Empty()
}

type MyStack struct {
	arr []int
}

func MyStackConstructor() MyStack {
	return MyStack{
		arr: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.arr = append(this.arr, x)
}

func (this *MyStack) Pop() int {
	if len(this.arr) > 0 {
		x := this.arr[len(this.arr)-1]
		this.arr = this.arr[:len(this.arr)-1]
		return x
	}
	return 0
}

func (this *MyStack) Top() int {
	if len(this.arr) > 0 {
		x := this.arr[len(this.arr)-1]
		return x
	}
	return 0
}

func (this *MyStack) Empty() bool {
	return len(this.arr) == 0
}
