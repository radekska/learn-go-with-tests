package arrays_slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 15, Sum(numbers))
	})

}

func TestSumAll(t *testing.T) {
	t.Run("sums the single slice", func(t *testing.T) {
		assert.Equal(t, []int{50}, SumAll([]int{10, 20, 20}))
	})

	t.Run("sums the multiple slice", func(t *testing.T) {
		sliceOne := []int{10, 20, 20}
		sliceTwo := []int{5, 10, 15, 20}
		sliceThree := []int{8, 12}
		assert.Equal(t, []int{50, 50, 20}, SumAll(sliceOne, sliceTwo, sliceThree))
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sums the single slice", func(t *testing.T) {
		assert.Equal(t, []int{40}, SumAllTails([]int{10, 20, 20}))
	})

	t.Run("sums the multiple slice", func(t *testing.T) {
		sliceOne := []int{10, 20, 20}
		sliceTwo := []int{5, 10, 15, 20}
		sliceThree := []int{8, 12}
		assert.Equal(t, []int{40, 45, 12}, SumAllTails(sliceOne, sliceTwo, sliceThree))
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assert.Equal(t, want, got)
	})
}
