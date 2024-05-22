package main

import "container/heap"

type elem struct {
	num, count int
}

type MinHeap []elem

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(elem))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	minHeap := &MinHeap{}
	counts := make(map[int]int)
	heap.Init(minHeap)

	for _, num := range nums {
		counts[num]++
	}

	for num, count := range counts {
		heap.Push(minHeap, elem{num, count})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	ans := make([]int, 0, k)

	for minHeap.Len() > 0 {
		ans = append(ans, heap.Pop(minHeap).(elem).num)
	}

	return ans
}
