/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type stackItem struct{
	node *TreeNode
	visit bool
}
func postorderTraversal(root *TreeNode) []int {
    stack := []stackItem{{root,false}}
	res := []int{}

	for len(stack) > 0{
		item := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if item.node != nil{
			if item.visit{
				res = append(res,item.node.Val)
			}else{
				stack = append(stack,stackItem{item.node,true})
				stack = append(stack,stackItem{item.node.Right,false})
				stack = append(stack,stackItem{item.node.Left,false})
			}
		}
	}
	return res
}
