package main

import "fmt"

func maxProduct(nums []int) int {

	size := len(nums)

	maxProduct := 1
	curProduct := 1

	for i := 0; i < size; i++ {
		curProduct *= nums[i]
		if curProduct > maxProduct {
			maxProduct = curProduct
		}
	}

	return maxProduct

}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
