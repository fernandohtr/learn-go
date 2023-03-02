package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	received := Perimeter(10.0, 10.0)
	expected := 40.0

	if received != expected {
		t.Errorf("\nreceived: %.2f\nexpected: %.2f", received, expected)
	}
}
