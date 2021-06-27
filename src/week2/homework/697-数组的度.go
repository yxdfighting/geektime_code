package homework

/*
   记录匹配到度的数字的左右边界，有可能有多个数字重复
*/

func findShortestSubArray(nums []int) int {
	ansMap := make(map[int]Helper)
	var maxD int
	res := 9999999999

	for idx, num := range nums {
		temp, ok := ansMap[num]
		if !ok {
			temp = Helper{}
		}

		temp.du++
		//更新左边界,第一次遇到后续不再更新
		if !temp.flag {
			temp.l = idx
			temp.flag = true
		}
		temp.r = idx
		ansMap[num] = temp
	}

	for _, v := range ansMap {
		if v.du == maxD && v.r-v.l+1 < res {
			res = v.r - v.l + 1
		}
		if v.du > maxD {
			// fmt.Printf("k,v:%v,%v\n",k,v)
			maxD = v.du
			res = v.r - v.l + 1
		}
	}

	return res
}

//定义每个数的左右边界，是否遇到第一次遇到该数fla为true，du表示该数的度
type Helper struct {
	l    int
	r    int
	du   int
	flag bool
}
