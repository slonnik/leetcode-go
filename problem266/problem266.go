package problem266

func canPermutePalindrome(s string) bool {
	data := make(map[rune]int)
	for _, symbol := range s {
		count, _ := data[symbol]
		data[symbol] = count + 1
	}

	for key, value := range data {
		if value%2 == 0 {
			delete(data, key)
		}
	}
	return len(data) < 2
}
