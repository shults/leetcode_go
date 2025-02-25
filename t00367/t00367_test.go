package t00367

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var res = isPerfectSquare(16)
	fmt.Printf("res=%+v", res)
}
