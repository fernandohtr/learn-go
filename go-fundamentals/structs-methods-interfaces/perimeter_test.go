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
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		received := rectangle.Area()
		expected := 72.0

		if received != expected {
			t.Errorf("\nreceived: %g\nexpected: %g", received, expected)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		received := circle.Area()
		expected := 314.1592653589793

		if received != expected {
			t.Errorf("\nreceived: %g\nexpected: %g", received, expected)
		}
	})
}
