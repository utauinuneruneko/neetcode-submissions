/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	res := -1
	cnt := 0
    var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if res >= 0 || cnt >= k {
			return
		}
		if node == nil {
			return
		}

		traverse(node.Left)
		cnt++
		if cnt == k {
			res = node.Val
			return
		}
		traverse(node.Right)
	}
	traverse(root)
	return res
}
