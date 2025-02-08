package main

import "github.com/stretchr/testify/require"

func Test_MaxNumberOfApples(t *testing.T) {
	require.Equal(t, MaxNumberOfApples([]int{100, 200, 150, 1000}), 1)
	require.Equal(t, MaxNumberOfApples([]int{900, 950, 800, 1000, 700, 800}), 2)
}
