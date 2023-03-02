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

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		received := shape.Area()
		if received != expected {
			t.Errorf("\nreceived: %g\nexpected: %g", received, expected)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
