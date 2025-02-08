package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSubarrayProductLessThanK(t *testing.T) {
	require.Equal(t, numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100), 8)
	require.Equal(t, numSubarrayProductLessThanK([]int{1, 2, 3}, 0), 0)
}
