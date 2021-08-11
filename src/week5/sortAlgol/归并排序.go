package sortAlgol

import "fmt"

/*
	分治思想，对nums分段进行排序，每次对分的两段进行合并
*/

func MergeSort(nums []int) {
	l, r := 0, len(nums)-1
	//nums只有一个元素时，就没必须要继续分割然后归并了
	if r-l < 1 {
		return
	}
	mid := (l + r) / 2

	MergeSort(nums[l : mid+1])
	MergeSort(nums[mid+1 : r+1])
	merge(nums, l, mid, r)
}

//mid算在前面
func merge(nums []int, l, mid, r int) {
	result := make([]int, r-l+1)
	left, right := l, mid+1
	fmt.Printf("nums:%v\n", nums)
	for i := 0; i < len(result); i++ {
		if left > mid {
			result[i] = nums[right]
			right++
			continue
		}
		if right > r {
			result[i] = nums[left]
			left++
			continue
		}

		if nums[left] >= nums[right] {
			result[i] = nums[right]
			right++
		} else {
			result[i] = nums[left]
			left++
		}
	}
	copy(nums, result)
	fmt.Printf("nums after:%v\n", nums)
}
