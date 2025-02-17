package t00046

import "slices"

func permute(nums []int) [][]int {
	switch {
	case len(nums) == 0:
		return [][]int{}
	case len(nums) == 1:
		return [][]int{{nums[0]}}
	case len(nums) == 2:
		return [][]int{
			{nums[0], nums[1]},
			{nums[1], nums[0]},
		}
	default:
		res := make([][]int, 0)

		for _, combination := range combinations(nums) {
			for _, cPerm := range permute(combination.rest) {
				res = append(res, slices.Concat([]int{combination.val}, cPerm))
			}
		}

		return res
	}
}

type comb struct {
	val  int
	rest []int
}

func combinations(nums []int) []comb {
	res := make([]comb, 0)

	for i, v := range nums {
		res = append(res, comb{
			val:  v,
			rest: slices.Concat(nums[:i], nums[i+1:]),
		})
	}

	return res
}
