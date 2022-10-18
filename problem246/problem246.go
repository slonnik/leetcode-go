package problem246

// https://leetcode.com/problems/strobogrammatic-number/
func isStrobogrammatic(num string) bool {

	var data = make(map[uint8]uint8)
	data['0'] = '0'
	data['1'] = '1'
	data['6'] = '9'
	data['8'] = '8'
	data['9'] = '6'

	len := len(num)
	index := 0
	for index < len {
		valOne := num[index]
		valTwo, ok := data[valOne]
		if !ok || valTwo != num[len-index-1] {
			return false
		}
		index++
	}
	return true
}
