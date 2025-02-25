package t00109

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var size = listLength(head)
	var tree = generateBst(0, size-1)

	var curr = head

	traverseInorderLtr(tree, func(node *TreeNode) {
		node.Val = curr.Val
		curr = curr.Next
	})

	return tree
}

func generateBst(start, end int) *TreeNode {
	var tn TreeNode

	if start > end {
		return nil
	}

	if start == end {
		return &tn
	}

	var mid = (start + end) / 2
	tn.Left = generateBst(start, mid-1)
	tn.Right = generateBst(mid+1, end)

	return &tn
}

func traverseInorderLtr(node *TreeNode, fn func(*TreeNode)) {
	if node == nil {
		return
	}

	traverseInorderLtr(node.Left, fn)
	fn(node)
	traverseInorderLtr(node.Right, fn)
}

func listLength(node *ListNode) int {
	var size int

	for node != nil {
		size++
		node = node.Next
	}

	return size
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
