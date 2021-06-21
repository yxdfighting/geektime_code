package week1

/*
	连续序列问题考虑前缀和
*/

func maxSubArray(nums []int) int {
	//要多一个长度，主要是a[l] + a[l+1] + ...+a[r-1] + a[r] = s[r] - s[l-1]
	//第l个元素到第r个元素的和 从1开始,一般为了方便处理，在nums前补0
	//这样就都从1开始
	preList := make([]int, len(nums)+1)
	minpreList := make([]int, len(nums)+1)
	nums = append([]int{0}, nums...)

	//注意这个循环，对于前缀和数组的第一个元素，默认是0，s[1]对应到a[0]
	for i := 1; i < len(nums); i++ {
		preList[i] = preList[i-1] + nums[i]
	}

	//同样的 minpreList[i]表示前i个数的最小值
	for i := 1; i < len(nums); i++ {
		if preList[i] < minpreList[i-1] {
			minpreList[i] = preList[i]
		} else {
			minpreList[i] = minpreList[i-1]
		}
	}

	max := -999999999

	// fmt.Printf("nums,preList,minPreList:%v,%v,%v",nums,preList,minpreList)

	//固定i,注意细节，j到i
	for i := 1; i < len(preList); i++ {
		//preList[i] - preList[j-1] 表示j到i的连续子数组和最大，其实也就是preList[j-1]最小
		//所以可以提前用一个smin表示preList的前缀最小数组
		if preList[i]-minpreList[i-1] >= max {
			max = preList[i] - minpreList[i-1]
		}
	}

	return max
}
