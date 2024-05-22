package main

import (
	"container/heap"
	"fmt"
)

// func maxSlidingWindow(nums []int, k int) []int {
// 	ans := []int{}

// 	for i := 0; i <= len(nums)-k; i++ {
// 		max := nums[i]
// 		for j := i + 1; j < i+k; j++ {
// 			if max < nums[j] {
// 				max = nums[j]
// 			}
// 		}
// 		ans = append(ans, max)
// 	}

// 	return ans
// }

type Pair struct {
	index, value int
}

type PairHeap []Pair

func (h PairHeap) Len() int { return len(h) }
func (h PairHeap) Less(i, j int) bool {
	return h[i].value > h[j].value || (h[i].value == h[j].value && h[i].index < h[j].index)
}
func (h PairHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	h := &PairHeap{}
	heap.Init(h) // khởi tạo heap rỗng

	for i := 0; i < k; i++ { // khởi tạo heap với k phần tử đầu tiên
		heap.Push(h, Pair{i, nums[i]})
	}

	ans = append(ans, (*h)[0].value) // phần tử đầu tiên của heap là max của window đầu tiên

	for i := k; i < len(nums); i++ { // duyệt từ k đến hết mảng
		heap.Push(h, Pair{i, nums[i]}) // thêm phần tử mới vào heap
		for (*h)[0].index <= i-k {     // nếu phần tử đầu tiên của heap không còn trong window hiện tại thì xóa nó đi
			heap.Pop(h) // xóa phần tử đầu tiên của heap
		}
		ans = append(ans, (*h)[0].value) // phần tử đầu tiên của heap là max của window hiện tại
	}

	return ans
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
