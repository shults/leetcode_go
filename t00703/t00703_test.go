package t00703

import (
	"testing"
)

func TestName(t *testing.T) {
	var val = Constructor(3, []int{4, 5, 8, 2})
	val.Add(3)
	val.Add(5)
	val.Add(10)
	val.Add(9)
	val.Add(4)
}
