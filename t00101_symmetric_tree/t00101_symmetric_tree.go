package t00101_symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type relativePosition int

const (
	rLeft  relativePosition = 1
	rRight relativePosition = 2
)

type flatNode struct {
	val int
	rp  relativePosition
}

func isSymmetric(root *TreeNode) bool {
	left := make([]flatNode, 0)

	traverseLtr(root.Left, rLeft, func(i int, rp relativePosition) {
		left = append(left, flatNode{
			val: i,
			rp:  rp,
		})
	})

	right := make([]flatNode, 0)

	traverseRtl(root.Right, rRight, func(i int, rp relativePosition) {
		right = append(right, flatNode{
			val: i,
			rp:  rp,
		})
	})

	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if (left[i].val != right[i].val) || (left[i].rp == right[i].rp) {
			return false
		}
	}

	return true
}

func traverseLtr(node *TreeNode, rp relativePosition, fn func(int, relativePosition)) {
	if node == nil {
		return
	}

	traverseLtr(node.Left, rLeft, fn)
	fn(node.Val, rp)
	traverseLtr(node.Right, rRight, fn)
}

func traverseRtl(node *TreeNode, rp relativePosition, fn func(int, relativePosition)) {
	if node == nil {
		return
	}

	traverseRtl(node.Right, rRight, fn)
	fn(node.Val, rp)
	traverseRtl(node.Left, rLeft, fn)
}
