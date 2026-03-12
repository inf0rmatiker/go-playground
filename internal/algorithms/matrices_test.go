package algorithms

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyMatricesNaive(t *testing.T) {
	a := [][]int{
		{3, 1, 4},
	}

	b := [][]int{
		{4, 3},
		{2, 5},
		{6, 8},
	}

	expected := [][]int{
		{38, 46},
	}

	actual := multiply2DNaive(a, b)

	assert.True(t, slices.EqualFunc(expected, actual, func(a, b []int) bool {
		return slices.Equal(a, b)
	}))
	fmt.Println(actual)
}

func TestMultiplyMatricesConcurrent(t *testing.T) {
	a := [][]int{
		{3, 1, 4},
		{7, 2, 3},
	}

	b := [][]int{
		{4, 3},
		{2, 5},
		{6, 8},
	}

	expected := [][]int{
		{38, 46},
		{50, 55},
	}

	actual := multiply2DConcurrent(a, b)

	assert.True(t, slices.EqualFunc(expected, actual, func(a, b []int) bool {
		return slices.Equal(a, b)
	}))
	fmt.Println(actual)
}

func TestMultiplyMatricesConcurrentByRow(t *testing.T) {
	a := [][]int{
		{3, 1, 4},
		{7, 2, 3},
	}

	b := [][]int{
		{4, 3},
		{2, 5},
		{6, 8},
	}

	expected := [][]int{
		{38, 46},
		{50, 55},
	}

	actual := multiply2DConcurrentByRow(a, b)

	assert.True(t, slices.EqualFunc(expected, actual, func(a, b []int) bool {
		return slices.Equal(a, b)
	}))
	fmt.Println(actual)
}
