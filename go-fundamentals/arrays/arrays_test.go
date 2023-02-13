package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	received := Sum(numbers)
	expected := 15

	if received != expected {
		t.Errorf("\nreceived: %d\nexpected: %d\ngiven numbers: %v", received, expected, numbers)
	}
}
