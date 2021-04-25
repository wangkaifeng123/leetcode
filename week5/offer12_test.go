package week5

import "testing"

//[["A","B","C","E"],["S","F","E","S"],["A","D","E","E"]]
//"ABCESEEEFS"
func TestExist(t *testing.T) {
	var insert [][]byte
	insert = append(insert, []byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'E', 'S'}, []byte{'A', 'D', 'E', 'E'})
	t.Log(exist(insert, "ABCESEEEFS"))
}

func exist(board [][]byte, word string) bool {
	boards := make([][]int, len(board))
	var res bool
	for i := 0; i < len(board); i++ {
		boards[i] = make([]int, len(board[0]))
	}
	for k, v := range board {
		for i := 0; i < len(v); i++ {
			if v[i] == word[0] {
				//先找到第一个 四方遍历 已经找到过的 不找
				boards[k][i] = 1
				res = helper(k, i, word, 0, board, boards)
				boards[k][i] = 0
				if res == true {
					return true
				}
			}
		}
	}
	return res
}

func helper(line, column int, word string, i int, board [][]byte, boards [][]int) bool {
	//更新序列
	if i == len(word)-1 {
		return true
	}
	if line-1 >= 0 {
		if boards[line-1][column] != 1 && board[line-1][column] == word[i+1] {
			boards[line-1][column] = 1
			if res := helper(line-1, column, word, i+1, board, boards); res == true {
				return true
			}
			boards[line-1][column] = 0
		}
	}
	if column+1 < len(board[0]) {
		if boards[line][column+1] != 1 && board[line][column+1] == word[i+1] {
			boards[line][column+1] = 1
			if res := helper(line, column+1, word, i+1, board, boards); res == true {
				return res
			}
			boards[line][column+1] = 0
		}
	}

	if column-1 >= 0 {
		if boards[line][column-1] != 1 && board[line][column-1] == word[i+1] {
			boards[line][column-1] = 1
			if res := helper(line, column-1, word, i+1, board, boards); res == true {
				return true
			}
			boards[line][column-1] = 0
		}
	}

	if line+1 < len(board) {
		if boards[line+1][column] != 1 && board[line+1][column] == word[i+1] {
			boards[line+1][column] = 1
			if res := helper(line+1, column, word, i+1, board, boards); res == true {
				return res
			}
			boards[line+1][column] = 0
		}
	}

	return false
}
