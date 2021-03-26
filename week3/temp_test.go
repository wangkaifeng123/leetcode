package week3

import "testing"

func TestHead(t *testing.T) {
	t.Log(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	heap := make([]int, 0) //用来存下标
	for k, v := range T {
		//推入栈中 从栈顶开始找 推出元素
		if len(heap) == 0 {
			heap = append(heap, k)
			continue
		}
		for i := len(heap) - 1; i >= 0; i-- {
			//当大于栈顶元素
			if T[heap[i]] < v {
				res[heap[i]] = k - heap[i]
				heap = heap[:len(heap)-1]
				if len(heap) == 0 {
					heap = append(heap, k)
				}
				continue //继续找下一个
			}
			heap = append(heap, k)
			break
		}
	}
	if len(heap) != 0 {
		for _, v := range heap {
			res[v] = 0
		}
	}
	return res
}
