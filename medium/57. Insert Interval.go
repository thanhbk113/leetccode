package main

import "sort"

func insert57(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	return merge(intervals)
}

func isOverlap(a, b []int) bool {
	return a[0] <= b[1] && b[0] <= a[1]
}

func mergeOverlap(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func merge57(intervals [][]int) [][]int {
	results := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		a := intervals[i]
		b := results[len(results)-1]
		if isOverlap(a, b) {
			results[len(results)-1] = mergeOverlap(a, b)
		} else {
			results = append(results, a)
		}
	}

	return results
}
