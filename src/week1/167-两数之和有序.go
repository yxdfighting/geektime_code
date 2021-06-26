package week1

/*
   因为有序，所以左下标向右移动时，右下标应该对应移动
*/

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	var res []int
	for l < r {
		if numbers[l]+numbers[r] == target {
			res = append(res, l+1)
			res = append(res, r+1)
			return res
		}

		if numbers[l]+numbers[r] > target {
			r--
		}

		if numbers[l]+numbers[r] < target {
			l++
		}
	}
	return res
}
