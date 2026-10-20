/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	cnt := 0
    var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, maxNum int) {
		if node == nil {
			 return
		}
		if maxNum <= node.Val {
			cnt++
		}
		traverse(node.Left, max(maxNum, node.Val))
		traverse(node.Right, max(maxNum, node.Val))
	}
	traverse(root, -101)
	return cnt
}
