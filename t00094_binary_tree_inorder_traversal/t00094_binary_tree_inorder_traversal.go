package t00094_binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	traverse(root, func(i int) {
		res = append(res, i)
	})

	return res
}

func traverse(root *TreeNode, fn func(int)) {
	if root == nil {
		return
	}

	traverse(root.Left, fn)
	fn(root.Val)
	traverse(root.Right, fn)
}
