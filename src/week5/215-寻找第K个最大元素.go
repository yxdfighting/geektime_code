package week5

import "math/rand"

/*
	采用快排思想，判断每次的pivot，如果k在pivot之后，就只排后半段

*/
func findKthLargest(nums []int, k int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	pivot := partition(nums)

	if pivot >= k-1 {
		//说明是在前半段
		return findKthLargest(nums[:pivot+1], k)
	} else {
		return findKthLargest(nums[pivot+1:], k-pivot-1)
	}
}

func partition(nums []int) int {
	randIdx := rand.Intn(len(nums))
	randVal := nums[randIdx]

	l, r := 0, len(nums)-1

	for l <= r {
		for nums[l] > randVal {
			l++
		}
		for nums[r] < randVal {
			r--
		}

		if l == r {
			break
		}
		if l < r {
			nums[r], nums[l] = nums[l], nums[r]
			r--
			l++
		}
	}
	return r
}
