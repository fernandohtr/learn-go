package integers

import "fmt"
import "testing"

func TestAdder(t *testing.T) {
	received := Add(2, 2)
	expected := 4

	if received != expected {
		t.Errorf("\nreceived: %d\nexpected: %d", received, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
