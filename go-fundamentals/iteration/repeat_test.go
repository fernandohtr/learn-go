package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	received := Repeat("a", 5)
	expected := "aaaaa"

	if received != expected {
		t.Errorf("\nreceived: %q\nexpected: %q", received, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
