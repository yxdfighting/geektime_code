package week2

/*
   类似上题，本题是子集题结果的子集
*/

var preRes []int
var ans [][]int

func combine(n int, k int) [][]int {
	preRes = []int{}
	ans = [][]int{}
	findCombine(1, k, n)
	return ans
}

func findCombine(level int, k int, n int) {
	//递归结束条件,已经选了k个数  或者选到n
	if len(preRes) == k {
		temp := make([]int, len(preRes))
		copy(temp, preRes)
		ans = append(ans, temp)
		return
	}

	if level == n+1 {
		return
	}

	//选取第level个数
	preRes = append(preRes, level)
	findCombine(level+1, k, n)
	preRes = preRes[0 : len(preRes)-1]
	//不选去第level个数
	findCombine(level+1, k, n)
}
