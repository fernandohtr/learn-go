package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		have := Hello("Fernando")
		want := "Hello, Fernando"
		assertCorrectMessage(t, have, want)
	})

	t.Run("say 'Hello, World' when an ampty string is supplied", func(t *testing.T) {
		have := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, have, want)
	})
}

func assertCorrectMessage(t testing.TB, have, want string) {
	t.Helper()
	if have != want {
		t.Errorf("\nhave %q\nwant %q", have, want)
	}
}
