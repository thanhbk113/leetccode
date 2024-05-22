package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	x, y int
	dst  float64
}

func kClosest(points [][]int, k int) [][]int {
	pointsWithDst := make([]Point, len(points))
	for i, p := range points {
		x, y := p[0], p[1]
		dst := calcDst(x, y)
		pointsWithDst[i] = Point{x, y, dst}
	}

	quickSelect(pointsWithDst, 0, len(pointsWithDst)-1, k) // find the k closest points

	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		ans[i] = []int{pointsWithDst[i].x, pointsWithDst[i].y}
	}

	return ans
}

func calcDst(x, y int) float64 {
	return math.Sqrt(float64(x*x + y*y))
}

func quickSelect(points []Point, left, right, k int) {
	if left < right {
		pivotIndex := partition(points, left, right) // partition
		if pivotIndex == k {                         // if pivotIndex is equal to k that means we have found the k closest points
			return
		} else if pivotIndex < k { // if pivotIndex is less than k that means we need to look for the k closest points in the right partition
			quickSelect(points, pivotIndex+1, right, k)
		} else { // if pivotIndex is greater than k that means we need to look for the k closest points in the left partition
			quickSelect(points, left, pivotIndex-1, k)
		}
	}
}

func partition(points []Point, left, right int) int {
	rand.Seed(time.Now().UnixNano())                                      // seed random number generator
	pivotIndex := left + rand.Intn(right-left+1)                          // select a random pivot index
	points[pivotIndex], points[right] = points[right], points[pivotIndex] // swap the pivot with the last element

	pivot := points[right]
	i := left
	for j := left; j < right; j++ {
		if points[j].dst < pivot.dst {
			points[i], points[j] = points[j], points[i] // swap the element at index i with the element at index j to place the element at index j at its correct position
			i++
		}
	}
	points[i], points[right] = points[right], points[i] // swap the pivot with the element at index i to place the pivot at its correct position
	return i                                            // return the index of the pivot
}

func main() {
	//test cases
	rand.Seed(time.Now().UnixNano())
	left := 0
	right := 2
	pivotIndex := left + (right-left)/2
	fmt.Println(pivotIndex)
}
