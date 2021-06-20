package week1

/*
	考虑通过栈来实现，栈的特点最近相关性，括号问题就是典型的最近相关性
	大体思路：遇到左括号都入栈，遇到右括号时看是否能与栈顶抵消,注意每次判断是否为空栈

	golang里面没有栈的数据结构
*/

func isValid(s string) bool {
	var byteStack []string

	for _, sByte := range s {
		if string(sByte) == "(" || string(sByte) == "[" || string(sByte) == "{" {
			byteStack = append(byteStack, string(sByte))
		}

		if string(sByte) == ")" {
			if len(byteStack) == 0 {
				byteStack = append(byteStack, string(sByte))
				continue
			}
			if byteStack[len(byteStack)-1] == "(" {
				byteStack = byteStack[0 : len(byteStack)-1]
			} else {
				byteStack = append(byteStack, string(sByte))
			}
		}

		if string(sByte) == "]" {
			if len(byteStack) == 0 {
				byteStack = append(byteStack, string(sByte))
				continue
			}
			if byteStack[len(byteStack)-1] == "[" {
				byteStack = byteStack[0 : len(byteStack)-1]
			} else {
				byteStack = append(byteStack, string(sByte))
			}
		}

		if string(sByte) == "}" {
			if len(byteStack) == 0 {
				byteStack = append(byteStack, string(sByte))
				continue
			}
			if byteStack[len(byteStack)-1] == "{" {
				byteStack = byteStack[0 : len(byteStack)-1]
			} else {
				byteStack = append(byteStack, string(sByte))
			}
		}
	}

	if len(byteStack) == 0 {
		return true
	}

	return false
}
