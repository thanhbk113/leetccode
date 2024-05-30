package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for _, coin := range coins {
		for x := coin; x <= amount; x++ {
			dp[x] = min(dp[x], dp[x-coin]+1) // x - coin là số tiền còn lại sau khi thêm một đồng coin
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{2}
	amount := 11
	fmt.Println(coinChange(coins, amount)) // 3
}
