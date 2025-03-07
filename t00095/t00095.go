package t00095

func generateTrees(n int) []*TreeNode {
	var cache = make(map[int][]*TreeNode)
	return generateTreesAux(0, n, cache)
}

func mapNode(node *TreeNode, diff int) *TreeNode {
	if node == nil {
		return nil
	}

	var mapped = new(TreeNode)
	mapped.Val = node.Val + diff
	mapped.Left = mapNode(node.Left, diff)
	mapped.Right = mapNode(node.Right, diff)

	return mapped
}

func mapPermutations(in []*TreeNode, diff int) []*TreeNode {
	var res = make([]*TreeNode, len(in))

	for i, v := range in {
		res[i] = mapNode(v, diff)
	}

	return res
}

func generateTreesAux(base int, n int, cache map[int][]*TreeNode) []*TreeNode {
	n = n - base

	if c, ok := cache[n]; ok {
		return mapPermutations(c, base)
	}

	var res = make([]*TreeNode, 0)

	switch n {
	case 0:
	case 1:
		var node = new(TreeNode)
		node.Val = 1
		res = append(res, node)
	default:

		var leftLen = 0
		var rightLen = n - 1

		for leftLen < n {
			var val = leftLen + 1
			var leftPermutations = generateTreesAux(0, leftLen, cache)
			var rightPermutations = generateTreesAux(rightLen, n-1, cache)

			if len(leftPermutations) == 0 {
				leftPermutations = append(leftPermutations, nil)
			}

			if len(rightPermutations) == 0 {
				rightPermutations = append(rightPermutations, nil)
			}

			for _, left := range leftPermutations {
				for _, right := range rightPermutations {
					var root = new(TreeNode)
					root.Val = val
					root.Left = left
					root.Right = right

					res = append(res, root)
				}
			}

			leftLen++
			rightLen--
		}

		cache[n] = res
		return res

	}

	cache[n] = res
	return res
}

type rng struct {
	start int
	end   int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
