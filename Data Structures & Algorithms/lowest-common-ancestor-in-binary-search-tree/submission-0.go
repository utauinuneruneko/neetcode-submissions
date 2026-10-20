/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    var find func(*TreeNode, *TreeNode, *TreeNode) *TreeNode
	find = func(node *TreeNode, a *TreeNode, b *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val == a.Val || node.Val == b.Val {
			return node
		}
		left, right := find(node.Left, a, b), find(node.Right, a, b)
		if left != nil && right != nil {
			return node
		} else if left != nil {
			return left
		} else {
			return right
		}
	}
	return find(root, p, q)
}
