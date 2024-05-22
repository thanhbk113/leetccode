package main

import (
	"fmt"
	"sort"
	"strconv"
)

func subsetsWithDup(nums []int) [][]int {
	checking := make(map[string]struct{})
	ans := [][]int{}

	sort.Ints(nums)

	backtrackWithDup(&ans, []int{}, 0, nums, checking)

	return ans
}

func getStringKey(curr []int) string {
	key := ""
	for _, val := range curr {
		key += strconv.Itoa(val)
	}

	return key
}

func backtrackWithDup(ans *[][]int, curr []int, start int, nums []int, checking map[string]struct{}) {
	tmp := make([]int, len(curr))
	copy(tmp, curr)
	*ans = append(*ans, tmp)

	for i := start; i < len(nums); i++ {
		curr = append(curr, nums[i])
		var key = getStringKey(curr)
		if _, ok := checking[key]; ok {
			curr = curr[:len(curr)-1]
			continue
		}
		checking[key] = struct{}{}
		backtrackWithDup(ans, curr, i+1, nums, checking)
		curr = curr[:len(curr)-1]
	}
}

func main() {
	nums := []int{1, 1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

// Input: nums = [1,2,2]
// Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

// -------------------------------------------
// nums = [1,2,2]

// []

// 1 - 2 = [1,2]
//   - 2 if exist - continues

// 2 - 2 = [2,2]
