package t00069_sqrtx

const MAX_X = (1 << 32) - 1
const APPROX_NR_OF_ELEMENTS = 1 << 16

var solutionsTable []int

func getSolutionsTable() []int {
	if solutionsTable == nil {
		solutionsTable = make([]int, 0, APPROX_NR_OF_ELEMENTS)

		for i := 0; i*i < MAX_X; i++ {
			solutionsTable = append(solutionsTable, i*i)
		}
	}

	return solutionsTable
}

func mySqrt(x int) int {
	srcTable := getSolutionsTable()

	return binSearch(
		srcTable,
		x,
		0,
		len(srcTable)-1,
	)
}

func binSearch(
	srcTable []int,
	x int,
	start int,
	end int,
) int {
	if end-start == 1 {
		return start
	}

	middle := start + (end-start)/2

	if x < srcTable[middle] {
		return binSearch(
			srcTable,
			x,
			start,
			middle,
		)
	} else {
		return binSearch(
			srcTable,
			x,
			middle,
			end,
		)
	}
}
