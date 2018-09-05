package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append([]int{math.MinInt64}, append(nums1, math.MaxInt64)...)
	nums2 = append([]int{math.MinInt64}, append(nums2, math.MaxInt64)...)
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	lenSum := len(nums1) + len(nums2)
	imin := 0
	imax := len(nums1) - 1
	i := int((imin + imax) / 2)
	j := int((len(nums1)+len(nums2)-4)/2) - i
	for (i >= 0 && i < len(nums1)-1 && j >= 0) &&
		!((nums1[i] <= nums2[j+1]) && (nums1[i+1] >= nums2[j])) {
		if nums1[i] > nums2[j+1] {
			imax = i
		} else if nums1[i+1] < nums2[j] {
			imin = i + 1
		}
		i = int((imin + imax) / 2)
		j = int((len(nums1)+len(nums2)-4)/2) - i
	}
	if (lenSum)%2 == 0 {
		return float64(max(nums1[i], nums2[j])+min(nums1[i+1], nums2[j+1])) / 2
	}
	return float64(min(nums1[i+1], nums2[j+1]))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
