package t00059

func generateMatrix(n int) [][]int {
	upTo := n * n
	vec := vector{1, 0}
	pos := position{0, 0}

	var res = make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for i := 1; i <= upTo; i++ {
		res[pos.x][pos.y] = i * i

		if !vec.hasNext(n, res, pos) {
			vec.rotate()
		}

		pos = pos.add(vec)
	}

	return res
}

type vector struct {
	x int
	y int
}

func (v *vector) rotate() {
	switch {
	case v.x == 1:
		v.x = 0
		v.y = -1
	case v.y == 1:
		v.x = -1
		v.y = 0
	case v.x == -1:
		v.x = 0
		v.y = -1
	case v.y == -1:
		v.x = 0
		v.y = -1
	}
}

func (v *vector) hasNext(size int, matrix [][]int, pos position) bool {
	pos = pos.add(*v)

	if pos.x < 0 || pos.x == size {
		return false
	}

	if pos.y < 0 || pos.y == size {
		return false
	}

	return matrix[pos.x][pos.y] > 0
}

type position struct {
	x int
	y int
}

func (p position) add(v vector) position {
	p.x += v.x
	p.y += v.y
	return p
}
