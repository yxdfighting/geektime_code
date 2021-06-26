package week1

/*
   用map，k为元素 v为出现的下标，因为题目说明同一个元素不能重复出现
*/
func twoSum(nums []int, target int) []int {
	var res []int
	numMap := make(map[int]int, len(nums))
	for idx, num := range nums {
		//如果在map中能找到，直接return，找不到就加入map
		if v, ok := numMap[target-num]; ok {
			res = append(res, idx)
			res = append(res, v)
			return res
		}
		numMap[num] = idx
	}

	return res
}
