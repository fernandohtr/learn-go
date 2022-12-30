package main

import "testing"

func TestHello(t *testing.T) {
	have := Hello("Fernando")
	// have := Hello()
	want := "Hello, Fernando"
	// want := "Hello, world"

	if have != want {
		t.Errorf("\nhave %q\nwant %q", have, want)
	}
}
