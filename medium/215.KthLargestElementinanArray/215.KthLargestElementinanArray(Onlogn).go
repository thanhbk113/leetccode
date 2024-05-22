package main

import (
	"container/heap"
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if k > n {
		return -1
	}

	h := &IntArr{}

	heap.Init(h)

	for _, val := range nums {
		heap.Push(h, val)
	}

	ans := 0
	for i := 0; i < k; i++ {
		ans = heap.Pop(h).(int)
	}

	return ans
}

type IntArr []int

func (h IntArr) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntArr) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntArr) Len() int {
	return len(h)
}

func (h *IntArr) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntArr) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	a := findKthLargest(nums, k)
	fmt.Println("🚀 ~ funcmain ~ a : ", a)
}
