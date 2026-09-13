/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    result := []int{}
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode){
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return result
}
