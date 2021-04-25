package week5

import "testing"

func TestCoinChange(t *testing.T) {
	t.Log(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {
	//动态转移方程式 f[0]=min(f[amount-coins[0],f[amount-coins[1],...)+1
	if amount == 0 {
		return 0
	}
	res := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		var condition []int
		for _, v := range coins {
			if i-v >= 0 {
				if res[i-v] == 0 && i-v != 0 {
					continue
				}
				condition = append(condition, res[i-v])
			}
		}
		if len(condition) == 0 {
			res[i] = 0
			continue
		}
		res[i] = min(condition)
	}
	if res[amount] == 0 {
		return -1
	}
	return res[amount]
}

func min(temp []int) int {
	res := temp[0]
	for _, v := range temp {
		if v <= res {
			res = v
			continue
		}
	}
	return res + 1
}
