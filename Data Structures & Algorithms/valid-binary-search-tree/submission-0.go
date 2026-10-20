/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    if root == nil {
		return true
	}

	var validate func(*TreeNode, int, int) bool
	validate = func(node *TreeNode, minVal, maxVal int) bool {
		if node == nil {
			return true
		}

		if node.Val <= minVal || node.Val >= maxVal {
			return false
		}

		return validate(node.Left, minVal, node.Val) && validate(node.Right, node.Val, maxVal)
	}

	return validate(root, math.MinInt32, math.MaxInt32)
}
