package _25_用队列实现栈

type MyStack struct {
	queue MyQueue
}

func Constructor() MyStack {
	return MyStack{
		queue: MyQueueConstructor(),
	}
}

func (this *MyStack) Push(x int) {
	this.queue.Push(x)
	n := this.queue.Size()
	for n > 1 {
		this.queue.Push(this.queue.Pop())
		n--
	}
}

func (this *MyStack) Pop() int {
	return this.queue.Pop()
}

func (this *MyStack) Top() int {
	return this.queue.Front()
}

func (this *MyStack) Empty() bool {
	return this.queue.Empty()
}

type MyQueue struct {
	arr []int
}

func MyQueueConstructor() MyQueue {
	return MyQueue{
		arr: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)
}

func (this *MyQueue) Pop() int {
	if len(this.arr) > 0 {
		x := this.arr[0]
		this.arr = this.arr[1:len(this.arr)]
		return x
	}
	return 0
}

func (this *MyQueue) Front() int {
	if len(this.arr) > 0 {
		x := this.arr[0]
		return x
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	return len(this.arr) == 0
}

func (this *MyQueue) Size() int {
	return len(this.arr)
}
