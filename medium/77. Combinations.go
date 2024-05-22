package main

import "fmt"

func combine(n int, k int) [][]int {
	var ans [][]int
	backtrack(&ans, []int{}, 1, n, k)

	return ans

}

func backtrack(res *[][]int, curr []int, start, n, k int) {
	if k == 0 {
		temp := make([]int, len(curr))
		copy(temp, curr)
		*res = append(*res, temp)
		return
	}

	for i := start; i <= n; i++ {
		curr = append(curr, i)
		backtrack(res, curr, i+1, n, k-1)

		curr = curr[:len(curr)-1]
	}
}

func main() {
	fmt.Println(combine(4, 2))
}

// n=4        k=2
// (1,2,3,4)  (1,2)

// i=1 (1->4) => [1,2],[1,3],[1,4] , k = 2 , curr = []{1}

// i=2 (2->4) => [2,3], [2,4]      , k = 1  curr =[]{2}

// i=3 (3 -> 4) => [3,4]                   , k =0   curr = []{3}
