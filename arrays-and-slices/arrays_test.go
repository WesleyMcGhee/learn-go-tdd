package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5", func (t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	}) 

	t.Run("Collection of 3", func (t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Testing 2 slices of 2", func (t *testing.T) {
		got := SumAll([][]int{[]int{1, 2}, []int{3, 4}})
		want := []int{3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Testing 3 slices of 2", func (t *testing.T) {
		got := SumAll([][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}})
		want := []int{3, 7, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Testing 2 slices of 3", func (t *testing.T) {
		got := SumAll([][]int{[]int{1, 2, 3}, []int{4, 5, 6}})
		want := []int{6, 15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}