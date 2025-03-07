package t00091

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var x = []int{1, 2, 3}
	fmt.Printf("res=%+v\n", x[0:3])

	//assert.Equal(t, 2, numDecodings("12"))
	//assert.Equal(t, 3, numDecodings("226"))

}
