package t00043

import (
	"fmt"
	"strings"
)

func multiply(num1 string, num2 string) string {
	a := parseBigint(num1)
	b := parseBigint(num2)
	return a.Mul(b).String()
}

type bigint struct {
	items []int
}

func parseBigint(num string) *bigint {
	nums := []byte(num)
	val := newBigint()

	for i, v := range nums {
		val.set(len(nums)-i-1, int(v-'0'))
	}

	return val
}

func newBigint() *bigint {
	return &bigint{
		items: make([]int, 1),
	}
}

func (b *bigint) String() string {
	var sb strings.Builder

	for i := len(b.items) - 1; i > -1; i-- {
		sb.WriteByte(byte(b.get(i)) + '0')
	}

	return sb.String()
}

func (b *bigint) Descr() string {
	return fmt.Sprintf("BigInt{%s}", b.String())
}

func (b *bigint) Len() int {
	return len(b.items)
}

func (b *bigint) Mul(other *bigint) *bigint {
	res := newBigint()

	bigger, lower := b, other

	if bigger.Len() < lower.Len() {
		bigger, lower = lower, bigger
	}

	for i := 0; i < lower.Len(); i++ {
		base := bigger.leftShift(i)
		base.mulByDigit(lower.get(i))
		res = res.Add(base)
	}

	return res
}

func (b *bigint) Add(other *bigint) *bigint {
	result := newBigint()

	upto := max(b.Len(), other.Len())
	rest := 0

	for i := 0; i < upto; i++ {
		aa := b.get(i)
		bb := other.get(i)

		val := rest + aa + bb

		if val > 9 {
			rest = val / 10
			result.set(i, val%10)
		} else {
			rest = 0
			result.set(i, val)
		}
	}

	if rest > 0 {
		result.set(upto, rest)
	}

	return result
}

func (b *bigint) mulByDigit(digit int) {
	if digit < 0 || digit > 9 {
		panic("unexpected value")
	}

	if digit == 0 {
		b.items = []int{0}
		return
	}

	rest := 0

	for i := len(b.items) - 1; i >= 0; i-- {
		val := rest + b.items[i]*digit

		if val > 9 {
			rest = val / 10
			b.items[i] = val % 10
		} else {
			rest = 0
			b.items[i] = val
		}
	}

	if rest > 0 {
		b.items = append([]int{rest}, b.items...)
	}
}

func (b *bigint) leftShift(num int) *bigint {
	items := make([]int, len(b.items)+num)

	for i, val := range b.items {
		items[i] = val
	}

	return &bigint{
		items: items,
	}
}

func (b *bigint) set(offset int, digit int) {
	if digit < 0 || digit > 9 {
		panic("unexpected digit in set")
	}

	if offset >= len(b.items) {
		b.items = append(make([]int, offset-len(b.items)+1), b.items...)
	}

	b.items[len(b.items)-offset-1] = digit
}

func (b *bigint) get(offset int) int {
	if offset >= len(b.items) {
		return 0
	}

	return b.items[len(b.items)-offset-1]
}
