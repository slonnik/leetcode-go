package problem243

import "math"

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	var minDistance = math.MaxInt64
	var currentIndex = -1

	for index, value := range wordsDict {
		if value == word1 {
			if currentIndex != -1 && wordsDict[currentIndex] != word1 {
				shift := index - currentIndex
				if shift < minDistance {
					minDistance = shift
				}
			}
			currentIndex = index

		} else if value == word2 {
			if currentIndex != -1 && wordsDict[currentIndex] != word2 {
				shift := index - currentIndex
				if shift < minDistance {
					minDistance = shift
				}
			}
			currentIndex = index
		}
	}

	return minDistance
}
