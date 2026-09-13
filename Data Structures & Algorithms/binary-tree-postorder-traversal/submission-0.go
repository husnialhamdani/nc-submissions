/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil{
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		result = append(result, node.Val)
	}
	postOrder(root)
	return result
}
