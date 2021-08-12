package week5

import "math/rand"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	//根据第一个元素和第二个元素对二维数组进行排序
	//如果第一个元素相等，则按照第二个元素进行排序
	fastSort(intervals)
	//fmt.Printf("intervals:%v\n",intervals)
	var result [][]int

	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		//fmt.Printf("i,left,right,result:%v,%v,%v,%v\n",i,left,right,result)
		if right >= intervals[i][0] {
			right = max(intervals[i][1], right)
		} else {
			tmp := []int{left, right}
			result = append(result, tmp)
			left = intervals[i][0]
			right = intervals[i][1]
		}
		if i == len(intervals)-1 {
			result = append(result, []int{left, right})
		}
		//fmt.Printf("after left,right,result:%v,%v,%v\n",left,right,result)
	}

	return result
}

func fastSort(nums [][]int) {
	if len(nums) <= 1 {
		return
	}
	pivot := partition(nums)
	//fmt.Printf("pivot:%v\n",pivot)
	fastSort(nums[:pivot+1])
	fastSort(nums[pivot+1:])
}

func partition(nums [][]int) int {
	randIdx := rand.Intn(len(nums))
	val := nums[randIdx]

	l, r := 0, len(nums)-1
	//fmt.Printf("randIdx:%v\n",randIdx)
	for l <= r {
		//fmt.Printf("l,r:%v,%v\n",l,r)
		for less(nums[l], val) {
			l++
		}
		for less(val, nums[r]) {
			r--
		}
		if l == r {
			break
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return r
}

//前者小于后者 true
func less(nums1 []int, nums2 []int) bool {
	if nums1[0] == nums2[0] {
		return nums1[1] < nums2[1]
	}

	return nums1[0] < nums2[0]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
