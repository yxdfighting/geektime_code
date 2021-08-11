package sortAlgol

/*
	选择：从未排序的部分选最小值放到已经排序的末尾
	所以很容易想到对其的优化，就是堆排序，每次从未排序的部分取最小值，即构建小顶堆
*/
func SelectSort(nums []int) {
	//每次从未排序部分找到最小值后，与当前要放置的位置元素进行交换,这样就可以不用新开辟空间
	//i表示要未排序部分元素最小值要放置的位置
	for i := 0; i < len(nums); i++ {
		//对于每个i都要找未排序部分的最小值,同时要记录min出现的idx，方便后续交换
		min := nums[i]
		minIdx := i
		//j用来遍历未排序部分
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				minIdx = j
			}
		}
		//将最小值放在i的位置，将i交换到最小值的位置，用于后续比较
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}
