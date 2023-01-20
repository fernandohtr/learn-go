package integers

import "testing"

func TestAdder(t *testing.T) {
	received := Add(2, 2)
	expected := 4

	if received != expected {
		t.Errorf("\nreceived: %d\nexpected: %d", received, expected)
	}
}
