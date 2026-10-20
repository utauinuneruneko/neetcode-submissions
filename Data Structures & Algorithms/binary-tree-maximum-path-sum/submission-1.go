/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
    var maxPath func(*TreeNode) int
	maxPath = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := maxPath(node.Left), maxPath(node.Right)
		res = max(res, max(left+node.Val+right, max(left+node.Val, max(right+node.Val, node.Val))))
		return max(left+node.Val, max(right+node.Val, node.Val))
	}
	maxPath(root)
	return res
}
