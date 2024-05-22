package main

import (
	"fmt"
	"sort"
)

//Time complexity : sort.Slice O(nlogn) + O(n) = O(nlogn) ( using merge sort)

//Brute Force O(n^2)
// func merge(intervals [][]int) [][]int {

// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] < intervals[j][0]
// 	})

// 	n := len(intervals)

// 	for i := 0; i < n-1; i++ {
// 		if intervals[i][1] >= intervals[i+1][0] {
// 			if intervals[i][1] < intervals[i+1][1] {
// 				intervals[i][1] = intervals[i+1][1]
// 			}
// 			intervals = append(intervals[:i+1], intervals[i+2:]...)
// 			n--
// 			i--
// 		}
// 	}

// 	return intervals
// }

// sweep line O(nlogn)
func merge(intervals [][]int) [][]int {

	n := len(intervals)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < n; i++ {
		start, end := intervals[i][0], intervals[i][1]

		lastReVal := result[len(result)-1][1]

		if start <= lastReVal {
			if end > lastReVal {
				result[len(result)-1][1] = end
			}
		} else {
			result = append(result, []int{start, end})
		}
	}

	return result
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
