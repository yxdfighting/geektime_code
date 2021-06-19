package homework

/*
	主体思路：
	从后往前遍历，如果不等于9,+1后break；如果等于9，就置为0

*/
func plusOne(digits []int) []int {
	i := len(digits) - 1
	for i >= 0 {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
		i--
	}
	//处理特殊情况
	if i == -1 {
		res := make([]int, len(digits)+1)
		res[0] = 1
		return res
	}

	return digits
}
