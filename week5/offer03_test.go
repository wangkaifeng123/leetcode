package week5

import "testing"

func TestFindRepeatNumber(t *testing.T) {
	t.Log(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		if ok, _ := m[v]; ok {
			return v
		}
		m[v] = true
	}
	return 0
}
