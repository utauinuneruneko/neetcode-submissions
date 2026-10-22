/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	visited := make(map[int]*Node)
	var dfs func(*Node) *Node
	dfs = func(p *Node) *Node {
		if p == nil {
			return nil
		}
		if _, ok := visited[p.Val]; ok {
			return visited[p.Val]
		}
		
		
		neighbors := make([]*Node, len(p.Neighbors))
		cloneP := &Node{
			Val: p.Val,
			Neighbors: neighbors,
		}
		visited[p.Val] = cloneP
		
		for i, next := range p.Neighbors {
			cloneP.Neighbors[i] = dfs(next)
		}
		return cloneP
	}
	res := dfs(node)
	return res
}
