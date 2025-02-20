package t00144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 0, 100)

	push(&stack, root)

	res := make([]int, 0, 100)

	for {
		node := pop(&stack)

		if node == nil {
			return res
		}

		res = append(res, node.Val)

		if node.Right != nil {
			push(&stack, node.Right)
		}

		if node.Left != nil {
			push(&stack, node.Left)
		}
	}
}

func pop(stack *[]*TreeNode) *TreeNode {
	if len(*stack) == 0 {
		return nil
	}

	last := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return last
}

func push(stack *[]*TreeNode, node *TreeNode) {
	*stack = append(*stack, node)
}
