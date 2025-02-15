package t_0001_zero_zoom

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	gen := NewCaseGenerator([]int{59, 324, 915, 608, 779, 958, 814, 387, 454, 505})
	// 59 + 915 + 779 + 814 + 387
	// 324 + 608 + 958 + 505 + 454
	c := gen.findCases(3072)
	fmt.Printf("c=%s\n", c)
}

func NewCaseGenerator(cards []int) *CaseGenerator {
	return &CaseGenerator{
		cards: cards,
	}
}

type CaseGenerator struct {
	cards []int
}

func (r *CaseGenerator) findCases(target int) []string {
	nrOfCases := 1 << len(r.cards)
	var cases = make([]string, 0)

	for i := 0; i < nrOfCases; i++ {
		if r.getSumForCase(i) == target {
			cases = append(cases, fmt.Sprintf("%010b", i))
		}
	}

	return cases
}

func (r *CaseGenerator) getSumForCase(val int) int {
	sum := 0
	start := 0
	end := len(r.cards) - 1

	for offset := 0; offset < len(r.cards); offset++ {
		left := (1 & (val >> offset)) > 0

		if left {
			if offset%2 == 0 {
				sum += r.cards[start]
			}

			start++
		} else {
			if offset%2 == 0 {
				sum += r.cards[end]
			}
			end--
		}
	}

	return sum
}
