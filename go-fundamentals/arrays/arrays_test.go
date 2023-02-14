package arrays

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		received := Sum(numbers)
		expected := 6

		assertCorrectMessage(t, received, expected, numbers)
	})
}

func TestSumAll(t *testing.T) {
	
	received := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(received, expected) {
		t.Errorf("\nreceived: %v\nexpected: %v", received, expected)
	}
}

func assertCorrectMessage(t testing.TB, received int, expected int, numbers []int) {
	t.Helper()
	if received != expected {
		t.Errorf("\nreceived: %d\nexpected: %d\ngiven numbers: %v", received, expected, numbers)
	}
}
