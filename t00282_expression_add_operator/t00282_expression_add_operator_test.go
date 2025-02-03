package t00282_expression_add_operator

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Combinations(t *testing.T) {
	combs := Combinations("1234567890")

	r := make(map[int]int)
	for _, c := range combs {
		r[strings.Count(c, "_")] += 1
	}

	fmt.Printf("res123=%+v\n", combs)
	fmt.Printf("res123=%+v\n", len(combs))
	fmt.Printf("m=%+v\n", r)
	fmt.Printf("m1=%+v\n", OperatorCombinations(1))
	fmt.Printf("m2=%+v\n", OperatorCombinations(2))
	fmt.Printf("m3=%+v\n", OperatorCombinations(3))
}

func Test_addOperators2(t *testing.T) {
	//combs := addOperators("123", 6)
	combs := addOperators("6", 6)

	fmt.Printf("combs=%+v\n", combs)
	fmt.Printf("combs2=%+v\n", Combinations("6"))
}
