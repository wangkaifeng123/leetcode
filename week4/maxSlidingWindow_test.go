package week4

import "testing"

func TestMaxSlidingWindow(t *testing.T) {
	t.Log(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	if len(nums) == 0 {
		return res
	}
	dqueue := make([]int, 0)

	dqueue = append(dqueue, nums[0])
	for i := 1; i < k; i++ {
		dqueue = queueHandle(dqueue, nums[i])
	}

	res = append(res, dqueue[0])
	for i := k; i < len(nums); i++ {
		//插入元素
		dqueue = queueHandle(dqueue, nums[i])
		if nums[i-k] == dqueue[0] {
			dqueue = dqueue[1:]
		}
		res = append(res, dqueue[0])
	}
	return res
}

func queueHandle(dqueue []int, val int) []int {
	for k, v := range dqueue {
		if val > v {
			dqueue = dqueue[:k]
			break
		}
	}
	dqueue = append(dqueue, val)

	return dqueue
}
