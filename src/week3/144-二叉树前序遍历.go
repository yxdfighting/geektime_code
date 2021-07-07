package week3

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

var res []int

func preorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	findPreorder(root)
	return res
}

func findPreorder(root *TreeNode) {
	if root == nil {
		return
	}

	res = append(res, root.Val)
	findPreorder(root.Left)
	findPreorder(root.Right)
}
