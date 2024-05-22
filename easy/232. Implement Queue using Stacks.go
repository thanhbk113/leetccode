package main

type MyQueue struct {
	stackPush, stackPop []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stackPush = append(this.stackPush, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stackPop) != 0 {
		x := this.stackPop[len(this.stackPop)-1]
		this.stackPop = this.stackPop[:len(this.stackPop)-1]
		return x
	}
	for len(this.stackPush) > 0 {
		x := this.stackPush[len(this.stackPush)-1]
		this.stackPush = this.stackPush[:len(this.stackPush)-1]
		this.stackPop = append(this.stackPop, x)
	}

	x := this.stackPop[len(this.stackPop)-1]
	this.stackPop = this.stackPop[:len(this.stackPop)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.stackPop) != 0 {
		return this.stackPop[len(this.stackPop)-1]
	}
	for len(this.stackPush) > 0 {
		x := this.stackPush[len(this.stackPush)-1]
		this.stackPush = this.stackPush[:len(this.stackPush)-1]
		this.stackPop = append(this.stackPop, x)
	}

	return this.stackPop[len(this.stackPop)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stackPop) == 0 && len(this.stackPush) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
