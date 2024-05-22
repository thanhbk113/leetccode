package main

import "fmt"

func subsets(nums []int) [][]int {
	ans := [][]int{}

	backtrack(&ans, []int{}, 0, nums)

	return ans
}

func backtrack(ans *[][]int, curr []int, start int, nums []int) {
	temp := make([]int, len(curr))
	copy(temp, curr)
	*ans = append(*ans, temp)

	for i := start; i < len(nums); i++ {
		curr = append(curr, nums[i])
		backtrack(ans, curr, i+1, nums)

		curr = curr[:len(curr)-1]
	}
}

func main() {
	nums := []int{1, 2}
	fmt.Println(subsets(nums))
}
