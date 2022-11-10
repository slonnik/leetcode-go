package problem344

func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}
	start := 0
	end := len(s) - 1
	for {
		if start >= end {
			return
		}
		temp := s[start]
		s[start] = s[end]
		s[end] = temp
		start++
		end--
	}
}
