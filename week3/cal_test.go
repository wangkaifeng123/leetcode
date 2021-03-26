package week3

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	s = "+" + s
	stack := make([]int, 0)
	for cur := 0; cur != len(s); {
		next := searchOperators(s, cur+1)
		sub := strings.TrimSpace(s[cur+1 : next])
		num, _ := strconv.Atoi(sub)
		if s[cur] == '-' {
			num *= -1
		} else if s[cur] == '*' {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num *= pop
		} else if s[cur] == '/' {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num = pop / num
		}
		stack = append(stack, num)
		cur = next
	}
	sum := 0
	for _, num := range stack {
		sum += num
	}
	return sum
}

func searchOperators(s string, start int) int {
	for i := start; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			return i
		}
	}
	return len(s)
}
