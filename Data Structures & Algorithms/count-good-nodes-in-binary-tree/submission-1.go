/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Pair struct{
	node *TreeNode
	maxVal int
}
func goodNodes(root *TreeNode) int {
    count := 0
	if root == nil{
		return count
	}
	q := []Pair{{node : root,maxVal: root.Val}}
	for len(q) > 0{
		curr := q[0]
		q = q[1:]
		if curr.node.Val >= curr.maxVal{
			count++
		}
		newMax := curr.maxVal
		newMax = max(curr.maxVal,curr.node.Val)
		if curr.node.Left != nil{
			q = append(q,Pair{node: curr.node.Left,maxVal : newMax})
		}
		if curr.node.Right != nil{
			q = append(q,Pair{node: curr.node.Right,maxVal : newMax})
		}
	}
	return count
}
