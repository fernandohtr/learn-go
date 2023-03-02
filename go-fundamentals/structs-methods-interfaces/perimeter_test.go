package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	received := Perimeter(10.0, 10.0)
	expected := 40.0

	if received != expected {
		t.Errorf("\nreceived: %.2f\nexpected: %.2f", received, expected)
	}
}

func TestAr(t *testing.T) {
	received := Area(12.0, 6.0)
	expected := 72.0

	if received != expected {
		t.Errorf("\nreceived: %.2f\nexpected: %.2f", received, expected)
	}
}
