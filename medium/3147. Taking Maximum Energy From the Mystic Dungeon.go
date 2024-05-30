package main

import "math"

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	dp := make([]int, n+1)
	ans := math.MinInt
	for i := n - 1; i >= 0; i-- {
		dp[i] = energy[i]
		if i+k < n {
			dp[i] += dp[i+k]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
