package problem252

import (
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	data := []int{}

	for _, value := range intervals {
		data = append(data, value...)
	}
	sort.Ints(data)
	for len(data) > 0 {
		internal := data[:2]
		data = data[2:]
		contains := false
		for _, value := range intervals {
			if value[0] == internal[0] && value[1] == internal[1] {
				contains = true
				break
			}
		}
		if !contains {
			return false
		}
	}
	return true
}
