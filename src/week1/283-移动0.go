package week1

/*
	在该题中，遇到非0时就放入新的下标n，此时不会覆盖到有效元素，只会覆盖掉元数据中的0
	但不影响，在最后遍历完之后补充上就可以了
*/


func moveZeroes(nums []int)  {
	//主体思路：遇到非0时放入新的下标,此时会覆盖，所以处理完之后，把后面替换为0
	n := 0
	for i := 0;i < len(nums);i++{
		if nums[i] != 0{
			nums[n] = nums[i]
			n++
		}
	}

	for n < len(nums){
		nums[n] = 0
		n++
	}
}
