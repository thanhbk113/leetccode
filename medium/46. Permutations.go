package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	ans := [][]int{}

	used := map[int]bool{}

	backtrackPermute(&ans, []int{}, nums, used)

	return ans
}

func backtrackPermute(ans *[][]int, curr []int, nums []int, used map[int]bool) {
	if len(curr) == len(nums) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*ans = append(*ans, tmp)
	}

	for _, val := range nums {
		if used[val] {
			continue
		}
		curr = append(curr, val)
		used[val] = true
		backtrackPermute(ans, curr, nums, used)
		curr = curr[:len(curr)-1]
		used[val] = false
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// 1   -   2   -   3   -   4
//             -   4   -   3

//         3   -   2   -   4
//                 4   -   2

//     -   4   -   2   -   3
//             -   3   -   2
