package problem350

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	index1 := 0
	index2 := 0
	for {
		if index1 == len(nums1) || index2 == len(nums2) {
			break
		}
		if nums1[index1] == nums2[index2] {
			result = append(result, nums1[index1])
			index1++
			index2++
		} else if nums1[index1] < nums2[index2] {
			index1++
		} else {
			index2++
		}
	}
	return result
}
