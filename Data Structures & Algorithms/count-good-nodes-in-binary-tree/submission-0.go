/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    res := 0
	if root == nil{
		return res
	}
	maxVal := root.Val
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode){
		if node == nil{
			return
		}
		prevMax := maxVal
		if node.Val >= maxVal{
			res++
			maxVal = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
		maxVal = prevMax
	}
	dfs(root)
	return res
}
