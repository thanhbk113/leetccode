package main

import "fmt"

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	monotonicStack := make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		if len(monotonicStack) == 0 {
			monotonicStack = append(monotonicStack, heights[i])
			ans[i] = 0
			continue
		}

		tmp := 0

		for len(monotonicStack) > 0 && monotonicStack[len(monotonicStack)-1] < heights[i] {
			tmp++
			monotonicStack = monotonicStack[:len(monotonicStack)-1]
		}

		if len(monotonicStack) > 0 {
			tmp++
		}

		monotonicStack = append(monotonicStack, heights[i])

		ans[i] = tmp
	}

	return ans

}

func main() {

	heights := []int{10, 6, 8, 5, 11, 9}
	fmt.Println(canSeePersonsCount(heights))
}
