package t00415

import (
	"slices"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	var adder = NewStringAdder(num1, num2)
	return adder.Add()
}

type StringAdder struct {
	num1, num2 string
	result     []int
	rest       int
}

func NewStringAdder(num1, num2 string) *StringAdder {
	return &StringAdder{
		num1:   num1,
		num2:   num2,
		result: make([]int, 0, max(len(num1), len(num2))+1),
	}
}

func (s *StringAdder) Add() string {
	s.result = s.result[:]
	s.rest = 0

	s.addCommon()
	s.addFirst()
	s.addSecond()

	if s.rest > 0 {
		s.registerRankBit(0)
	}

	slices.Reverse(s.result)

	var builder strings.Builder
	builder.Grow(len(s.result))

	for _, val := range s.result {
		builder.WriteByte(byte(val + '0'))
	}

	return builder.String()
}

func (s *StringAdder) registerRankBit(sum int) {
	sum += s.rest

	if sum > 9 {
		s.rest = 1
		s.result = append(s.result, sum%10)
	} else {
		s.rest = 0
		s.result = append(s.result, sum)
	}
}

func (s *StringAdder) addCommon() {
	for i, upto := 1, min(len(s.num1), len(s.num2)); i <= upto; i++ {
		s.registerRankBit(
			int(s.num1[len(s.num1)-i]-'0') + int(s.num2[len(s.num2)-i]-'0'),
		)
	}
}

func (s *StringAdder) addFirst() {
	for i := min(len(s.num1), len(s.num2)) + 1; i <= len(s.num1); i++ {
		s.registerRankBit(int(s.num1[len(s.num1)-i] - '0'))
	}
}

func (s *StringAdder) addSecond() {
	for i := min(len(s.num1), len(s.num2)) + 1; i <= len(s.num2); i++ {
		s.registerRankBit(int(s.num2[len(s.num2)-i] - '0'))
	}
}
