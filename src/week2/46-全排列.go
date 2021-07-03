package week2

/*
   大体思路： 记录每一个元素的使用记录，每次从未使用的元素列表里选一个元素加入到preRes
*/

var preRes []int
var ans [][]int
var used []bool

func permute(nums []int) [][]int {
	preRes = []int{}
	ans = [][]int{}
	used = make([]bool, len(nums))
	findPermute(nums, 0)
	return ans
}

//level表示结果中第level位的选择
func findPermute(nums []int, level int) {
	if level == len(nums) {
		temp := make([]int, len(preRes))
		copy(temp, preRes)
		ans = append(ans, temp)
		return
	}
	//结果的第level位，可以从used里选取未选用的元素，选用后对应标志位为true
	for idx, num := range nums {
		if !used[idx] {
			preRes = append(preRes, num)
			used[idx] = true

			findPermute(nums, level+1)

			preRes = preRes[0 : len(preRes)-1]
			used[idx] = false
		}
	}
}
