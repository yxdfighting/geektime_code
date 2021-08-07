package week4

/*
	1、最大值最小 是一个求解问题，此类问题可以转化为 在状态空间中对判定问题进行枚举
	2、对状态空间的枚举，如果结果能够满足规律，就可以使用二分进行时间复杂度降级
*/

func splitArray(nums []int, m int) int {
	//确定左右边界,左边界表示能够满足>=子数组的和最大值，如果要分为m个子数组，至少应该要与整个数组最大值相等
	//右边界 如果子数组只划分为一个 那么此时应该为整个数组的和
	max, sum := -1, 0
	for _, obj := range nums {
		if obj > max {
			max = obj
		}
		sum += obj
	}

	// fmt.Printf("result15:%v\n",isValid(nums,m,15))

	l, r := max, sum

	for l < r {
		// fmt.Printf("l,r:%v,%v\n",l,r)
		mid := (l + r) / 2
		//如果isValid为true，表示是在右边 00001111  满足条件的最小值
		if isValid(nums, m, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

//判断t是否能大于等于子数组和的最大值
func isValid(nums []int, m int, t int) bool {
	sum := 0
	count := 1
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > t {
			sum = nums[i]
			count++
		} else {
			sum += nums[i]
		}
	}

	return count <= m
}
