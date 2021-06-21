package week1

func numberOfSubarrays(nums []int, k int) int {
	//根据奇偶转化为0，1
	for idx, num := range nums {
		nums[idx] = num % 2
	}

	//因为前缀和数组第0个数为0，所以在前面补个0
	zero := []int{0}
	nums = append(zero, nums...)

	//得到前缀和数组
	preList := make([]int, len(nums))
	for idx := 1; idx < len(preList); idx++ {
		preList[idx] = preList[idx-1] + nums[idx]
	}
	//fmt.Printf("nums,preList:%v,%v",nums,preList)
	countRes := 0
	//在前缀和数组中，对于每个i，找到一个j，使得preList[i]-preList[j-1]=k i从1到最后 j从1到i
	//for i := 1;i < len(preList);i++{
	//	for j := 0;j < i;j++{
	//		if preList[i] - preList[j] == k{
	//			count++
	//		}
	//	}
	//}
	//减少时间复杂度,固定i，j需要等于某个值，用一个数组统计，该值出现次数
	//count[i] = m表示的是i这个值在数组中出现了m次
	count := make([]int, len(preList)-1)
	for i := 1; i < len(preList); i++ {
		count[preList[i-1]]++
	}

	//fmt.Printf("count:%v",count)

	for i := 0; i < len(preList); i++ {
		if preList[i]-k >= 0 {
			countRes += count[preList[i]-k]
		}
	}

	return countRes
}
