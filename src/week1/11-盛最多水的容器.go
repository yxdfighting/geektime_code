package week1

/*
   容器的盛水决定于短边，依次遍历每一个容器，求出其最大盛水量
   固定左边容器，依次从右边去找，如果r大于l的盛水量，此时已经达到其最大值
   接着移动l
*/
func maxArea(height []int) int {
	var result int

	for i := 0; i < len(height); i++ {
		result = max(findMaxOne(height, i), result)
		// fmt.Printf("i:%v,result:%v\n",i,result)
	}

	return result
}

//找到某个i固定时的最大值
func findMaxOne(height []int, i int) int {
	var area int
	for j := len(height) - 1; j > 0; j-- {
		if height[j] >= height[i] {
			area = max(height[i]*(j-i), area)
			// fmt.Printf("j:%v\n",j)
			break
		} else {
			area = max(height[j]*(j-i), area)
		}
	}
	return area
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
