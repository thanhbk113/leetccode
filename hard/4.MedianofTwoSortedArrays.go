func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j, nTotal := 0, 0, len(nums1)+len(nums2)
	midArrIndex := nTotal / 2
	nums3 := []int{}

	for len(nums3) <= midArrIndex {
		if i < len(nums1) && (j >= len(nums2) || nums1[i] < nums2[j]) {
			nums3 = append(nums3, nums1[i])
			i++
		} else if j < len(nums2) {
			nums3 = append(nums3, nums2[j])
			j++
		}
	}

	if nTotal%2 == 0 {
		return float64((nums3[midArrIndex-1] + nums3[midArrIndex])) / 2
	} else {
		return float64(nums3[midArrIndex])
	}
}