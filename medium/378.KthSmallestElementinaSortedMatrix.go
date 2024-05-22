package main

import (
	"container/heap"
	"fmt"
)

func kthSmallest(matrix [][]int, k int) int {
	nLen, mLen := len(matrix), len(matrix[0])
	h := pq{}

	heap.Init(&h)
	for i := 0; i < mLen; i++ {
		h = append(h, Elem{r: i, c: 0, val: matrix[i][0]})
	}

	var res int

	for i := 0; i < k; i++ {
		e := heap.Pop(&h).(Elem)

		if e.c < nLen-1 {
			heap.Push(&h, Elem{r: e.r, c: e.c + 1, val: matrix[e.r][e.c+1]})
		}

		res = e.val
	}

	return res
}

type Elem struct {
	r   int
	c   int
	val int
}

type pq []Elem

func (h pq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h pq) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h pq) Len() int {
	return len(h)
}

func (h *pq) Push(elem interface{}) {
	*h = append(*h, elem.(Elem))
}

func (h *pq) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func main() {
	matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	a := kthSmallest(matrix, 8)
	fmt.Println("🚀 ~ funcmain ~ a : ", a)
}
