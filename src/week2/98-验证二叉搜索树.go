package week2

func isValidBST(root *TreeNode) bool {
	res, _, _ := checkValid(root)
	return res
}

//每次除判断有效外 还要返回当前节点的最大最小值
func checkValid(node *TreeNode) (bool, int, int) {
	if node == nil {
		return true, -9999999999999, 99999999999999
	}
	//判断左右子树是否有效
	leftRes, leftMax, leftMin := checkValid(node.Left)
	rightRes, rightMax, rightMin := checkValid(node.Right)

	valid := leftRes && rightRes && node.Val > leftMax && node.Val < rightMin
	//判断节点是否大于左子树最大值 小于右子树最小值
	return valid, max(max(leftMax, rightMax), node.Val), min(min(leftMin, rightMin), node.Val)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
