package t00039

func combinationSum(candidates []int, target int) [][]int {
	var res = make([][]int, 0)
	var prefix = make([]int, 0, 100)

	combinationSumAex(prefix, candidates, target, &res)

	return res
}

func combinationSumAex(prefix []int, candidates []int, target int, rest *[][]int) {
	for i, val := range candidates {
		if target == val {
			var row = make([]int, len(prefix)+1)
			copy(row, prefix)
			row[len(prefix)] = val
			*rest = append(*rest, row)
		} else if val > target {
			continue
		} else {
			combinationSumAex(
				append(prefix, val),
				candidates[i:],
				target-val,
				rest,
			)
		}
	}
}
