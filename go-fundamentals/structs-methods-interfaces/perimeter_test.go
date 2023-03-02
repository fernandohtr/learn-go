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

	t.Run("table driven tests", func(t *testing.T) {
		areaTests := []struct {
			shape    Shape
			expected float64
		}{
			{shape: Rectangle{12, 6}, expected: 72.0},
			{shape: Circle{10}, expected: 314.1592653589793},
			{shape: Triangle{12, 6}, expected: 36.0},
		}

		for _, tt := range areaTests {
			received := tt.shape.Area()
			if received != tt.expected {
				t.Errorf("\nreceived: %g\nexpected: %g", received, tt.expected)
			}
		}
	})
}

func TestArea2(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, expected: 72.0},
		{name: "Circle", shape: Circle{10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			received := tt.shape.Area()
			if received != tt.expected {
				t.Errorf("\n%#v\nreceived: %g\nexpected: %g", tt.shape, received, tt.expected)
			}
		})
	}
}
