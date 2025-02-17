package t00119

import "math"

func getRow(rowIndex int) []int {
	switch rowIndex {
	case 0:
		return []int{1}
	case 1:
		return []int{1, 1}
	default:
		rowIndex += 1
		res := make([]int, rowIndex)
		res[0] = 1
		res[1] = 1
		res[len(res)-1] = 1

		for i := 2; i < rowIndex; i++ {
			upto := i / 2
			mid := int(math.Ceil(float64(i) / 2))

			for j := 0; j < upto; j++ {
				res[upto-j] = res[upto-j] + res[upto-j-1]
				res[mid+j] = res[upto-j]
			}
		}

		return res
	}
}
