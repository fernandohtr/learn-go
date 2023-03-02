package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	received := Perimeter(rectangle)
	expected := 40.0

	if received != expected {
		t.Errorf("\nreceived: %.2f\nexpected: %.2f", received, expected)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	received := Area(rectangle)
	expected := 72.0

	if received != expected {
		t.Errorf("\nreceived: %.2f\nexpected: %.2f", received, expected)
	}
}
