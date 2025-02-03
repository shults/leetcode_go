package t00069_sqrtx

const Min = 0
const Max = int(float64(1<<15) * 1.5)

func mySqrt(x int) int {
	start, end := Min, Max

	for {
		if end-start == 1 {
			return start
		}

		middle := start + (end-start)/2

		if x < middle*middle {
			end = middle
		} else {
			start = middle
		}
	}
}
