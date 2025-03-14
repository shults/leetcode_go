package t00566

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	var res = matrixReshape(
		[][]int{{1, 2}, {3, 4}},
		1,
		4,
	)

	fmt.Printf("%v\n", res)
}
