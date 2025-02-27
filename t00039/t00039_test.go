package t00039

import (
	"fmt"
	"testing"
)

func Test(T *testing.T) {
	var res = combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Printf("res=%+v\n", res)

	var res2 = combinationSum([]int{2, 3, 5}, 8)
	fmt.Printf("res2=%+v\n", res2)
}
