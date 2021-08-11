package sortAlgol

import (
	"fmt"
	"math/rand"
)

/*
	本质也是个分治思想，每次找一个基准，将数组交换为基准前的元素小于基准，基准后的元素大于基准
	从左开始，如果小于基准，则继续往右
	从右开始，如果大于基准，则继续往左
*/
func FastSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivot := partition(nums)

	FastSort(nums[:pivot+1])
	FastSort(nums[pivot+1:])
}

//随机找一个基准，分成前后两部分
//遇到等于，也要交换
func partition(nums []int) int {
	fmt.Printf("nums:%v\n", nums)
	l, r := 0, len(nums)-1
	//基准选取采用随机方式
	randIdx := rand.Intn(r + 1)
	val := nums[randIdx]
	fmt.Printf("randIdx:%v\n", randIdx)
	for l <= r {
		//如果遇到等于的情况就要停下来，与右边的遍历结果进行交换
		for nums[l] < val {
			l++
		}
		for nums[r] > val {
			r--
		}
		if l == r {
			break
		}
		if l < r {
			//此时交换l和r
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}

	}
	fmt.Printf("after sort:%v,r:%v\n", nums, r)
	return r
}
