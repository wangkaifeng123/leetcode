package week5

import "testing"

func TestMinimumTotal(t *testing.T) {
	var triangle [][]int
	//triangle=append(triangle,[]int{2},[]int{3,4},[]int{6,5,7},[]int{4,1,8,3})
	triangle = append(triangle, []int{-1}, []int{-2, -3})
	t.Log(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	var res int
	res = triangle[0][0]
	temp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		temp[i] = make([]int, i+1)
	}
	temp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		//遍历每一层的元素
		for j := 0; j <= i; j++ {
			if j == 0 {
				temp[i][j] = temp[i-1][j] + triangle[i][j]
				if i == len(triangle)-1 {
					res = temp[i][j]
				}
				continue
			}
			if j == i {
				temp[i][j] = temp[i-1][j-1] + triangle[i][j]
				if i == len(triangle)-1 {
					res = less(res, temp[i][j])
				}
				continue
			}
			temp[i][j] = less(temp[i-1][j], temp[i-1][j-1]) + triangle[i][j]
			if i == len(triangle)-1 {
				res = less(res, temp[i][j])
			}
		}
	}

	return res
}

func less(a, b int) int {
	if a > b {
		return b
	}
	return a
}
