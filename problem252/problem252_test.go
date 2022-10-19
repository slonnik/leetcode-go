package problem252

import "testing"

func TestCaseOne(t *testing.T) {
	var data [][]int = [][]int{[]int{0, 30}, []int{5, 10}, []int{15, 20}}
	result := canAttendMeetings(data)
	if result {
		t.Errorf("Can't attend meetings for %v", data)
	}
}

func TestCaseTwo(t *testing.T) {
	var data [][]int = [][]int{[]int{7, 10}, []int{2, 4}}
	result := canAttendMeetings(data)
	if !result {
		t.Errorf("Can attend meetings for %v", data)
	}
}
