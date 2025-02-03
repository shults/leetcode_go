package t00282_expression_add_operator

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"slices"
	"strconv"
	"strings"
)

// todo: finish cause of this case is not finished

const (
	Plus  Operator = '+'
	Minus Operator = '-'
	Mul   Operator = '*'
)

type Operator byte

type OperatorTuple []Operator

func addOperators(
	num string,
	target int,
) []string {
	combs := Combinations(num)
	res := make([]string, 0)

	for _, comb := range combs {
		nums := strings.Count(comb, "_")
		operatorCombs := OperatorCombinations(nums)

		for _, operatorComb := range operatorCombs {
			srcExp := merge(comb, operatorComb)
			expr, err := parser.ParseExpr(srcExp)

			if err != nil {
				panic(err)
			}

			evalRes, err := evalInt(expr)

			if evalRes == target {
				res = append(res, srcExp)
			}
		}

		if len(operatorCombs) == 0 {
			expr, err := parser.ParseExpr(comb)

			if err != nil {
				panic(err)
			}

			evalRes, err := evalInt(expr)

			if evalRes == target {
				res = append(res, comb)
			}
		}
	}

	return res
}

func merge(nums string, ops OperatorTuple) string {
	res := ""
	numbers := strings.Split(nums, "_")

	for i, operator := range ops {
		res += numbers[i] + string(operator)
	}

	res += numbers[len(numbers)-1]

	return res
}

func Combinations(num string) []string {
	cache := make(map[string]struct{})

	base := strings.Split(num, "")

	combinations(base, cache)

	res := make([]string, 0, len(cache))

	for k, _ := range cache {
		res = append(res, k)
	}

	return res
}

func combinations(
	base []string,
	cache map[string]struct{},
) {
	key := strings.Join(base, "_")
	cache[key] = struct{}{}

	for i := 0; i < len(base)-1; i++ {
		if base[i] == "0" {
			continue
		}

		combinations(
			slices.Concat(base[0:i], []string{base[i] + base[i+1]}, base[i+2:]),
			cache,
		)
	}
}

var operatorCache = make(map[int][]OperatorTuple)

func OperatorCombinations(size int) []OperatorTuple {
	return operatorCombinations(size, operatorCache)
}

func operatorCombinations(size int, cache map[int][]OperatorTuple) []OperatorTuple {
	res, ok := cache[size]

	if ok {
		return res
	}

	if size == 0 {
		res = []OperatorTuple{}
		cache[size] = res
		return res
	}

	if size == 1 {
		res = []OperatorTuple{
			{Plus},
			{Minus},
			{Mul},
		}

		cache[size] = res

		return res
	}

	childSet := operatorCombinations(size-1, cache)
	res = make([]OperatorTuple, 0, len(childSet)*3)

	for _, op := range []Operator{
		Plus,
		Minus,
		Mul,
	} {
		for _, cs := range childSet {
			res = append(res, append(OperatorTuple{op}, cs...))
		}
	}

	return res
}

func evalInt(expr ast.Expr) (int, error) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		i, err := strconv.ParseInt(e.Value, 10, 64)

		if err != nil {
			return 0, err
		}

		return int(i), nil
	case *ast.BinaryExpr:
		left, err := evalInt(e.X)

		if err != nil {
			return 0, err
		}

		right, err := evalInt(e.Y)

		if err != nil {
			return 0, err
		}

		switch e.Op {
		case token.MUL:
			return left * right, nil
		case token.ADD:
			return left + right, nil
		case token.SUB:
			return left - right, nil
		default:
			return 0, errors.New("unexpected operator")
		}
	default:
		return 0, errors.New("unexpected expression")
	}
}
