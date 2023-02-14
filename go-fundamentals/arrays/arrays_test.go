package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		received := Sum(numbers)
		expected := 6

		assertCorrectMessage(t, received, expected, numbers)
	})
}

func assertCorrectMessage(t testing.TB, received int, expected int, numbers []int) {
	t.Helper()
	if received != expected {
		t.Errorf("\nreceived: %d\nexpected: %d\ngiven numbers: %v", received, expected, numbers)
	}
}
