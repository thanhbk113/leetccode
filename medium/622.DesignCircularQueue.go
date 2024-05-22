package main

type MyCircularQueue struct {
	values   []int
	capacity int
	head     int
	tail     int
	size     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		values:   make([]int, k),
		capacity: k,
		head:     -1,
		tail:     -1,
		size:     0,
	}

}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = 0
	}
	this.tail = (this.tail + 1) % this.capacity
	this.values[this.tail] = value
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = (this.head + 1) % this.capacity
	this.size--
	if this.size == 0 {
		this.head = -1
		this.tail = -1
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.values[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.values[this.tail]

}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
