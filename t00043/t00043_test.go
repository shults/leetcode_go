package t00043

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	for _, tCase := range []struct {
		num1   string
		num2   string
		result string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
	} {
		t.Run(fmt.Sprintf("%s + %s = %s", tCase.num1, tCase.num2, tCase.result), func(t *testing.T) {
			assert.Equal(t, tCase.result, multiply(tCase.num1, tCase.num2))
			assert.Equal(t, tCase.result, multiply(tCase.num2, tCase.num1))
		})
	}
}

func TestParse(t *testing.T) {
	strInt := parseBigint("12345")
	assert.Equal(t, "BigInt{12345}", strInt.Descr())
}

func Test_leftShift(t *testing.T) {
	strInt := parseBigint("123")
	assert.Equal(t, "1230", strInt.leftShift(1).String())
	assert.Equal(t, "12300", strInt.leftShift(2).String())
}

func Test_add(t *testing.T) {
	assert.Equal(t, "1122", parseBigint("123").Add(parseBigint("999")).String())
	assert.Equal(t, "1998", parseBigint("999").Add(parseBigint("999")).String())
}
