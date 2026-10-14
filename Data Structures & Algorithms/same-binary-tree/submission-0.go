/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    return isEqual(p,q)
}

func isEqual(p *TreeNode, q *TreeNode) bool{
	if p == nil && q == nil{
		return true
	}
	if p == nil || q == nil{
		return false
	}

	if p.Val != q.Val{
		return false
	}

	return isEqual(p.Left,q.Left) && isEqual(p.Right,q.Right)
}