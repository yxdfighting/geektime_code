package homework

/*
	区间连续性问题，考虑前缀和 前缀和的下标表示前i个元素之和
	a[i] + a[i+1] + ....a[n] = s[n] - s[i-1]   i从0到n-1
*/

func subarraySum(nums []int, k int) int {
	//构造前缀和数组
	preList := make([]int, len(nums)+1)

	for i := 1; i < len(preList); i++ {
		preList[i] = preList[i-1] + nums[i-1]
	}
	//fmt.Printf("%v",preList)

	count := 0
	for i := 1; i < len(preList); i++ {
		//j表示区间的idx-1
		for j := 0; j < i; j++ {
			if preList[i]-preList[j] == k {
				count++
			}
		}
	}

	return count
}
