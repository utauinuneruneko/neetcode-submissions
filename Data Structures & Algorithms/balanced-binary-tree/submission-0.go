/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	flag := true
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if !flag || node == nil {
			return 0
		}
		left, right := depth(node.Left), depth(node.Right)
		if int(math.Abs(float64(left-right))) > 1 {
			flag = false
		}
		return max(left, right) + 1
	}
	depth(root)
	return flag
}
