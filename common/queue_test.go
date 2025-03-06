package common_test

import (
	"github.com/stretchr/testify/require"
	"leetcode_go/common"
	"testing"
)

func Test(t *testing.T) {
	q := common.NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	val := q.Dequeue()
	require.Equal(t, 1, val)
	require.Equal(t, 2, q.Len())

	val = q.Dequeue()
	require.Equal(t, 2, val)
	require.Equal(t, 1, q.Len())

	val = q.Dequeue()
	require.Equal(t, 3, val)
	require.Equal(t, 0, q.Len())

}
