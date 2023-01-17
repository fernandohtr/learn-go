package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		received := Hello("Fernando", "")
		expected := "Hello, Fernando"
		assertCorrectMessage(t, received, expected)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		received := Hello("", "")
		expected := "Hello, world"
		assertCorrectMessage(t, received, expected)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		received := Hello("Fernando", "Portuguese")
		expected := "Olá, Fernando"
		assertCorrectMessage(t, received, expected)
	})
	t.Run("in Spanish", func(t *testing.T) {
		received := Hello("Fernando", "Spanish")
		expected := "Hola, Fernando"
		assertCorrectMessage(t, received, expected)
	})
}

func assertCorrectMessage(t testing.TB, received, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("\nreceived: %q\nexpected: %q", received, expected)
	}
}
