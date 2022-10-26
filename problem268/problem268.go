package problem268

func missingNumber(nums []int) int {

	min := nums[0]
	max := nums[0]
	sum := 0

	for _, item := range nums {
		if item < min {
			min = item
		}
		if item > max {
			max = item
		}
		sum = sum + item
	}

	expectedSum := (max - min + 1) * (max + min) / 2

	if expectedSum == sum {
		if min == 0 {
			return max + 1
		}
		return 0
	}

	return expectedSum - sum
}
