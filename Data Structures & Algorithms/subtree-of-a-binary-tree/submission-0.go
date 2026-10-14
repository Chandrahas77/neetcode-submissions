/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    //base case
	if root == nil{
		return false
	}
	if isSameTree(root,subRoot){
		return true
	}
	return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)
}

func isSameTree(tree, subTree *TreeNode) bool{
	if tree == nil && subTree == nil{
		return true
	}
	if tree == nil || subTree == nil{
		return false
	}
	if tree.Val != subTree.Val{
		return false
	}
	return isSameTree(tree.Left,subTree.Left) && isSameTree(tree.Right,subTree.Right)
}