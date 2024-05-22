
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.getLengthMinStack() == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		firstValMinStack := this.minStack[this.getLengthMinStack()-1]
		if val < firstValMinStack {
			this.minStack = append(this.minStack, val)
		} else {
			this.minStack[this.getLengthMinStack()-1] = firstValMinStack
			this.minStack = append(this.minStack, firstValMinStack)
		}
	}
}

func (this *MinStack) getLengthMinStack() int {
	return len(this.minStack)
}

func (this *MinStack) getLengthStack() int {
	return len(this.minStack)
}

func (this *MinStack) Pop() {
	if this.getLengthStack() == 0 {
		return
	}
	this.stack = this.stack[:this.getLengthStack()-1]
	this.minStack = this.minStack[:this.getLengthMinStack()-1]
}

func (this *MinStack) Top() int {
	if this.getLengthStack() == 0 {
		return -1
	}
	return this.stack[this.getLengthStack()-1]
}

func (this *MinStack) GetMin() int {
	if this.getLengthMinStack() == 0 {
		return -1
	}
	return this.minStack[this.getLengthMinStack()-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */