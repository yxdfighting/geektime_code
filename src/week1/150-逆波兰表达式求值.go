package week1

import (
	"strconv"
	"strings"
)

/*
	基本思路：遇到数字就入栈，遇到字符就将栈顶的两个数与符号进行计算，再将结果压栈
*/

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		if existInSlice([]string{"+", "-", "*", "/"}, token) {
			if len(stack) >= 2 {
				second := stack[len(stack)-1]
				first := stack[len(stack)-2]
				stack = stack[0 : len(stack)-2]
				stack = append(stack, compute(token, first, second))
			}

		} else {
			stack = append(stack, changer(token))
		}
	}

	return stack[len(stack)-1]
}

func existInSlice(result []string, token string) bool {
	for _, res := range result {
		if res == token {
			return true
		}
	}
	return false
}

func changer(token string) int {
	if strings.HasPrefix(token, "-") {
		res, _ := strconv.Atoi(token[1:len(token)])
		return 0 - res
	}
	res, _ := strconv.Atoi(token)
	return res
}

func compute(token string, first, second int) int {
	switch token {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	case "/":
		return first / second
	default:
		return 0
	}
}
