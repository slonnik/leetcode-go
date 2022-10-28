package problem283

func moveZeroes(nums []int) {
	len := len(nums)
	shift := 0
	for i := 0; i < len; i++ {
		if nums[i] == 0 {
			shift++
			continue
		}
		nums[i-shift] = nums[i]
		if shift > 0 {
			nums[i] = 0
		}
	}
}
