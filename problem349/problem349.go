package problem349

func intersection(nums1 []int, nums2 []int) []int {
	data := make([]int, 1001)

	for _, value := range nums1 {
		if data[value] == 0 {
			data[value] = 1
		}
	}

	for _, value := range nums2 {
		if data[value] == 1 {
			data[value] = 2
		}
	}

	result := []int{}
	startPos := 0
	endPos := len(data) - 1
	for {
		if startPos > endPos {
			break
		}

		if data[startPos] == 2 {
			result = append(result, startPos)
		}

		if startPos == endPos {
			break
		}
		if data[endPos] == 2 {
			result = append(result, endPos)
		}
		startPos++
		endPos--

	}
	return result
}
