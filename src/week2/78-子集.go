package week2

/*
   问题可以分解为 每个元素取或者不取 n个元素可以分为n个level，转化为一个树的模型

   1. 什么情况下需要新定义一个方法做递归？
   因为要递归n次才能有一个结果，所以level要作为一个参数进行传递，题目中的方法显然不符合条件，
   所以需要新定义一个

   2. 为什么每次需要还原现场？
   在该次递归中修改了全局的变量，想象树的模型，做完依次递归到n之后，这时候要从树的底层回溯，
   回溯的时候要把上次的改动还原
           1
       1,2    1
   比如上面，[1,2]递归完选取2的时候，就要往上回溯，不选取2，此时pre数组其实是[1,2]所以就要把2添加到pre里面的行为还原
   3. 为什么每次把pre放入到ans的时候要深拷贝一次？
   因为slice在golang中的实现是struct，放入ans中的其实是内存块的地址，后续每次修改pre依然修改已经放入ans的pre，因为其实是同一个地址
*/
//存每次的中间结果,
//为什么要定义全局？
//pre存放遍历n个元素之后的结果，要在多次递归调用中传递，所以是个全局的
//ans存在最终的结果，肯定也是全局的
var pre []int
var ans [][]int

func subsets(nums []int) [][]int {
	pre = pre[0:0]
	ans = ans[0:0]
	findSubsets(nums, 0)
	return ans
}
func findSubsets(nums []int, level int) {
	//终止条件，把n层都递归完
	if level == len(nums) {
		temp := make([]int, len(pre))
		copy(temp, pre)
		ans = append(ans, temp)
		return
	}
	//第level不选
	findSubsets(nums, level+1)
	//第level要选
	pre = append(pre, nums[level])
	findSubsets(nums, level+1)
	pre = pre[0 : len(pre)-1]
}
