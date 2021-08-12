package week5

//计数排序，先用一个数组表示i这个元素出现了res[i]次
func relativeSortArray(arr1 []int, arr2 []int) []int {
	// var resultPart1,resultPart2 []int
	var resultPart1 []int
	//先找到最大值
	max := 0
	for i := 0; i < len(arr1); i++ {
		if arr1[i] > max {
			max = arr1[i]
		}
	}
	times := make([]int, max+1)

	for i := 0; i < len(arr1); i++ {
		times[arr1[i]]++
	}

	// fmt.Printf("times:%v\n",times)

	for j := 0; j < len(arr2); j++ {
		for k := 0; k < times[arr2[j]]; k++ {
			resultPart1 = append(resultPart1, arr2[j])
		}
		times[arr2[j]] = 0
	}

	for arr, time := range times {
		if time != 0 {
			for k := 0; k < time; k++ {
				resultPart1 = append(resultPart1, arr)
			}
		}
	}
	return resultPart1
}
