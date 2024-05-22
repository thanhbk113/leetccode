package main

import "container/heap"

type Heap struct {
	Values   []int
	LessFunc func(int, int) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() int          { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(int)) }

func NewHeap(less func(int, int) bool) *Heap {
	return &Heap{LessFunc: less}
}

type MedianFinder struct {
	smallHeap *Heap
	largeHeap *Heap
}

func Constructor295() MedianFinder {
	return MedianFinder{
		smallHeap: NewHeap(func(a, b int) bool {
			return a < b
		}),
		largeHeap: NewHeap(func(a, b int) bool {
			return a > b
		}),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if (this.smallHeap.Len()+this.largeHeap.Len())%2 == 0 {
		heap.Push(this.smallHeap, num)
		heap.Push(this.largeHeap, heap.Pop(this.smallHeap))
	} else {
		heap.Push(this.largeHeap, num)
		heap.Push(this.smallHeap, heap.Pop(this.largeHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.smallHeap.Len()+this.largeHeap.Len())%2 == 0 {
		return (float64(this.smallHeap.Peek()) + float64(this.largeHeap.Peek())) / 2
	}
	return float64(this.largeHeap.Peek())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
