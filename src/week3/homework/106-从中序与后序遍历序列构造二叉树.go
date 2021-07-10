package homework

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
    首先考虑返回结果的根节点，继而类推到整个树
    1. 后序遍历 左右根 所以对应某棵树的后序遍历序列的最后一个元素一定是根
    2. 根据根元素将中序进行分割，前半段为左子树的中序遍历结果，后半段为右子树的中序遍历结果
    3. 根据2中分割出来的左右子树长度可以确定 对应左右子树的后续遍历 这样就找到了左右子树对应的中序和后序遍历结果 就找到了共同的子问题
    4. 问题归结到递归，递归终止条件，后序遍历的序列长度小于1时，这时候已经找不到该序列的根节点


一个比较重要的问题，什么情况下需要新定义一个方法用于递归？
*/

//注意go的默认初始化，对于结构体通过new或者 &Struct{}这种显示初始化的方式 会给一个对应类型的初始值
//var res *TreeNode 得到的就是一个nil

func buildTree(inorder []int, postorder []int) *TreeNode {
	var res *TreeNode
	//找root节点
	lenSlice := len(inorder)
	//终止递归,当序列中只有一个元素时，就不继续找左右子树了
	if lenSlice < 1 {
		return res
	}

	res = new(TreeNode)
	res.Val = postorder[lenSlice-1]
	if lenSlice == 1 {
		return res
	}

	//idx表示左序列中根节点的idx 遵循左闭右开
	var idx int
	//根据root节点在中序中找到左右子树的中序序列
	for idx, _ = range inorder {
		if inorder[idx] == res.Val {
			break
		}
	}

	//找到左子树
	res.Left = buildTree(inorder[:idx], postorder[:idx])
	if lenSlice > idx+1 {
		res.Right = buildTree(inorder[idx+1:], postorder[idx:lenSlice-1])
	}

	return res
}
