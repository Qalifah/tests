package sum

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{2,3,4,5,6}
		got := sum(nums)
		expected := 20

		if got != expected {
			t.Errorf("got %d expected %d given %#v", got, expected, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := sumAll([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 0})
	expected := []int{15, 30}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %d expected %d", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %d expected %d", got, expected)
		}
	}
	t.Run("make the sum of some slices", func(t *testing.T) {
		got := sumAllTails([]int{2, 3, 4, 5}, []int{7, 8, 9, 0})
		expected := []int{12, 17}
		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := sumAllTails([]int{}, []int{70, 89, 20})
		expected := []int{0, 109}
		checkSums(t, got, expected)
	})
}