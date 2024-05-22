package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]

	var less, greater []int

	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	less = quicksort(less)
	greater = quicksort(greater)

	return append(append(less, pivot), greater...)
}

func main() {
	arr := []int{7, 2, 1, 6, 8, 5, 3, 4}
	fmt.Println("Unsorted array:", arr)
	arr = quicksort(arr)
	fmt.Println("Sorted array:", arr)
}
