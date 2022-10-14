package problem242

func isAnagram(s string, t string) bool {
	var mapS = mapString(s)
	var mapT = mapString(t)
	if len(mapS) != len(mapT) {
		return false
	}

	for keyS, valueS := range mapS {
		valueT, ok := mapT[keyS]
		if !ok || valueS != valueT {
			return false
		}
	}

	return true
}

func mapString(s string) map[rune]int {
	var result = make(map[rune]int)
	for _, symbol := range s {
		value, _ := result[symbol]
		result[symbol] = value + 1

	}
	return result
}
