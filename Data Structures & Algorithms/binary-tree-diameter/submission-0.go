/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
    var maxDepth func(*TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := maxDepth(node.Left), maxDepth(node.Right)
		if left+right > res {
			res = left+right
		}
		return max(left, right)+1
	}
	maxDepth(root)
	return res
}
