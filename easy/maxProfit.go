package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minFt := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minFt {
			minFt = prices[i]
		} else if prices[i]-minFt > profit {

			profit = prices[i] - minFt
		}
	}

	return profit
}

func main() {
	// fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2}))
}
