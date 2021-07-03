package week2

var preRes []int
var ans [][]int

//表示在同一个位置已经放置过的元素
var used []bool

func permuteUnique(nums []int) [][]int {
	preRes = []int{}
	ans = [][]int{}
	used = make([]bool, len(nums))
	findPermuteUnique(nums, 0)
	return ans

}

//level表示第level个位置元素的放置
//在放置同一个位置时，如果放了一个元素，后续相同的元素就不能在这个位置再放了
func findPermuteUnique(nums []int, level int) {
	if level == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, preRes)
		ans = append(ans, temp)
		return
	}

	usedMap := make(map[int]bool, len(nums))

	for idx, num := range nums {
		if !usedMap[num] && !used[idx] {
			preRes = append(preRes, num)
			usedMap[num] = true
			used[idx] = true

			findPermuteUnique(nums, level+1)
			used[idx] = false
			preRes = preRes[0 : len(preRes)-1]
		}
	}
}
