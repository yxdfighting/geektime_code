package sortAlgol

//依次从头开始，遍历，每次遍历依次把最小值放在最前面
/*

 */
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
