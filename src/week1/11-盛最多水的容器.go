package week1

/*
   容器的盛水决定于短边，依次遍历每一个容器，求出其最大盛水量
   固定左边容器，依次从右边去找，如果r大于l的盛水量，此时已经达到其最大值
   接着移动l
*/

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0

	for i < j {
		// fmt.Printf("i,j:%v,%v\n",i,j)
		res = max(min(height[i], height[j])*(j-i), res)
		if height[i] == height[j] {
			i++
			j--
			continue
		}
		if height[i] > height[j] {
			j--
			continue
		}
		if height[i] < height[j] {
			i++
			continue
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
