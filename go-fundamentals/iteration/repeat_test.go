package iteration

import "testing"

func TestRepeat(t *testing.T) {
	received := Repeat("a")
	expected := "aaaaa"

	if received != expected {
		t.Errorf("\nreceived: %q\nexpected: %q", received, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
