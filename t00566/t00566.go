package t00566

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	var in = matrix(mat)
	var res = newMatrix(r, c)

	var top = r * c

	for i := 0; i < top; i++ {
		res.set(i, in.get(i))
	}

	return res
}

type matrix [][]int

func newMatrix(r, c int) matrix {
	var m = make([][]int, r)

	for i := 0; i < r; i++ {
		m[i] = make([]int, c)
	}

	return m
}

func (m matrix) rows() int {
	return len(m)
}

func (m matrix) cols() int {
	return len(m[0])
}

func (m matrix) size() int {
	return m.rows() * m.cols()
}

func (m matrix) indexes(i int) (c, r int) {
	r = i / (m.cols())
	c = i - (r * m.cols())
	return
}

func (m matrix) set(i, val int) {
	var c, r = m.indexes(i)
	m[r][c] = val
}

func (m matrix) get(i int) int {
	var c, r = m.indexes(i)
	return m[r][c]
}
