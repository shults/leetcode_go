package t00119

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Printf("%+v\n", getRow(i))
	}
}
