package demo3

//589. N叉树的前序遍历
//给定一个 N 叉树，返回其节点值的前序遍历。

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

//dfs
func Run(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, root.Val)
	if root.Children != nil {
		for _, v := range root.Children {
			res = append(res, Run(v)...)
		}
	}
	return res
}
