package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxLength(t *testing.T) {
	require.Equal(t, 5, maxLength([]int{1, 2, 1, 2, 1, 1, 1}))
	require.Equal(t, 3, maxLength([]int{2, 3, 4, 5, 6}))
	require.Equal(t, 5, maxLength([]int{1, 2, 3, 1, 4, 5, 1}))
	require.Equal(t, 3, maxLength([]int{1, 1, 9, 9, 10}))
}
