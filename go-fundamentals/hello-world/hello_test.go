package main

import "testing"

func TestHello(t *testing.T) {
	received := Hello()
	expected := "Hello, world"

	if received != expected {
		t.Errorf("\nreceived: %q\nexpected: %q", received, expected)
	}
}