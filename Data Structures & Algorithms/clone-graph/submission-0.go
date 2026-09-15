/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil{
		return nil
	}
    newNode := make(map[*Node]*Node)
	
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node{
		// if its already cloned
		if clone, exists := newNode[node]; exists{
			return clone
		}

		// create clone
		clone := &Node{
			Val: node.Val,
		}

		//store immediately
		newNode[node] = clone

		for _, neighbor := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
		}

		return clone
	}

	return dfs(node)
}
