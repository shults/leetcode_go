package t00103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result = make([][]int, 0, 1<<8)
	var queue = make([]Entry, 0, 2000)

	queue = append(queue, Entry{
		depth: 0,
		node:  root,
	})

	for len(queue) > 0 {
		var entry = queue[0]
		queue = queue[1:]

		if len(result) == entry.depth {
			result = append(result, make([]int, 0, 1<<entry.depth))
		}

		if entry.depth&1 == 0 {
			result[entry.depth] = append(result[entry.depth], entry.node.Val)
		} else {
			result[entry.depth] = append([]int{entry.node.Val}, result[entry.depth]...)
		}

		if entry.node.Left != nil {
			queue = append(queue, Entry{
				depth: entry.depth + 1,
				node:  entry.node.Left,
			})
		}

		if entry.node.Right != nil {
			queue = append(queue, Entry{
				depth: entry.depth + 1,
				node:  entry.node.Right,
			})
		}
	}

	return result
}

type Entry struct {
	depth int
	node  *TreeNode
}
