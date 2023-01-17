package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		received := Hello("Fernando")
		expected := "Hello, Fernando"
		assertCorrectMessage(t, received, expected)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		received := Hello("")
		expected := "Hello, world"
		assertCorrectMessage(t, received, expected)
	})
}

func assertCorrectMessage(t testing.TB, received, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("\nreceived: %q\nexpected: %q", received, expected)
	}
}
