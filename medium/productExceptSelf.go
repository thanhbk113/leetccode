package main

import "fmt"

func productExceptSelf(nums []int) []int {

	size := len(nums)

	left := make([]int, size)
	right := make([]int, size)

	left[0] = 1
	right[size-1] = 1

	for i := 1; i < size; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	for i := size - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < size; i++ {
		right[i] = left[i] * right[i]
	}

	return right

}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
