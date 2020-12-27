package demo5

//94. 二叉树的中序遍历
//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//输入：root = [1,null,2,3]
//输出：[1,3,2]

//首先我们需要了解什么是二叉树的中序遍历：按照访问左子树——根节点——右子树的方式遍历这棵树，
//而在访问左子树或者右子树的时候我们按照同样的方式遍历，直到遍历完整棵树。因此整个遍历过程天然具有递归的性质，我们可以直接用递归函数来模拟这一过程。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int


//方法一：递归
func Run(root *TreeNode) []int {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

//方法二：栈
func Run2(root *TreeNode) []int {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
