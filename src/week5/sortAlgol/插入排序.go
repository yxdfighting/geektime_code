package sortAlgol

//插入排序的主要思路可以参考打牌，每次拿到一张牌，要放在手中哪个位置
//每次插入已排序序列，就要发生移动，所以还是一个O(n2)
func InsertSort(nums []int) {
	//未排序区第一个元素依次往前遍历 比较
	//如果小于，则交换后继续往前；如果大于，就停止
	var j int
	for i := 1; i < len(nums); i++ {
		//j如果一直比前面的元素更小，就一直往前交换，直到走到最开始
		j = i
		for j >= 1 && nums[j] < nums[j-1] {
			nums[j], nums[j-1] = nums[j-1], nums[j]
			j--
		}
	}
}
