package week3

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res []int

func inorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	tempRes := findInorder(root)
	res = append(res, tempRes...)
	return res
}

func findInorder(root *TreeNode) []int {
	temp := make([]int, 0)

	if root == nil {
		return temp
	}

	temp = append(temp, findInorder(root.Left)...)
	temp = append(temp, root.Val)
	temp = append(temp, findInorder(root.Right)...)

	return temp
}
