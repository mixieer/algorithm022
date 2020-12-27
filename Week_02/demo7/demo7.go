package demo7

//429. N 叉树的层序遍历
//给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
//树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

//输入：root = [1,null,3,2,4,null,5,6]
//输出：[[1],[3,2,4],[5,6]]

type Node struct {
	Val      int
	Children []*Node
}

var res [][]int

//bfs
func Run(root *Node) [][]int {
	res = [][]int{}
	if root == nil {
		return res
	}

	queue := []*Node{root}
	var level int
	for 0 < len(queue) {
		counter := len(queue)
		res = append(res, []int{})
		for i := 0; i < counter; i++ {
			if queue[i] != nil {
				res[level] = append(res[level], queue[i].Val)
				for _, n := range queue[i].Children {
					queue = append(queue, n)
				}
			}
		}
		queue = queue[counter:]
		level++
	}
	return res
}
