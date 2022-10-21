package problem258

// https://leetcode.com/problems/add-digits/
func addDigits(num int) int {

	for num > 9 {
		var sum int
		for num > 9 {
			sum += num % 10
			num = num / 10
		}
		num += sum
	}
	return num
}
