package problem345

func reverseVowels(s string) string {

	if len(s) < 2 {
		return s
	}

	buffer := []byte(s)
	data := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

	startPos := 0
	endPos := len(buffer) - 1
	for {
		if startPos >= endPos {
			break
		}
		leftCharacter := buffer[startPos]
		rightCharacter := buffer[endPos]
		_, okLeft := data[leftCharacter]
		_, okRight := data[rightCharacter]

		if okLeft && okRight {
			buffer[startPos] = rightCharacter
			buffer[endPos] = leftCharacter
			startPos++
			endPos--
			continue
		}

		if !okLeft {
			startPos++
		}
		if !okRight {
			endPos--
		}
	}
	return string(buffer)
}
