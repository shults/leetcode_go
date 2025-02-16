package t00046

import (
	"fmt"
	"testing"
)

func Test_combinations(t *testing.T) {
	res := permute([]int{1, 2, 3})
	fmt.Printf("res=%v\n", res)
}

func Test_combinations_2(t *testing.T) {
	p := []int{1, 2, 3}
	fmt.Printf("%+v\n", p[:len(p)/2])
	fmt.Printf("%+v\n", p[len(p)/2:])

	p = []int{1, 2, 3, 4}
	fmt.Printf("%+v\n", p[:len(p)/2])
	fmt.Printf("%+v\n", p[len(p)/2:])
}
