package problem303

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	result := NumArray{data: make([]int, len(nums))}
	result.data[0] = nums[0]
	for index, value := range nums {
		if index == 0 {
			result.data[index] = value
		} else {
			result.data[index] = result.data[index-1] + value
		}

	}
	return result
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.data[right]
	}

	return this.data[right] - this.data[left-1]
}
