package arrays

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3}

	received := Sum(numbers)
	expected := 6

	assertCorrectMessage(t, received, expected, numbers)
}

func assertCorrectMessage(t testing.TB, received int, expected int, numbers []int) {
	t.Helper()
	if received != expected {
		t.Errorf("\nreceived: %d\nexpected: %d\ngiven numbers: %v", received, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	received := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(received, expected) {
		t.Errorf("\nreceived: %v\nexpected: %v", received, expected)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, received, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(received, expected) {
			t.Errorf("\nreceived: %v\nexpected: %v", received, expected)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		received := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, received, expected)
	})

	t.Run("safely sum empty slices", func (t *testing.T) {
		received := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, received, expected)
	})
}
