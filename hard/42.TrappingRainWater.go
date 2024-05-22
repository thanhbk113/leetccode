package main

func trap(h []int) int {
	n := len(h)
	waters := 0
	l, r := 0, n-1
	maxL, maxR := h[l], h[r]
	for l != r {
		if maxL <= maxR {
			waters += maxL - h[l]
			l++
			maxL = max(maxL, h[l])
		} else {
			waters += maxR - h[r]
			r--
			maxR = max(maxR, h[r])
		}
	}

	return waters
}
