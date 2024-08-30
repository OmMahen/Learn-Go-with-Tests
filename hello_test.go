package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Mahen", "")
		want := "Hello, Mahen"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World' when empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Ellodie", "Spanish")
		want := "Hola, Ellodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Mahen", "French")
		want := "Bonjour, Mahen"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // Mark this method as a helper so failures report the function call line, not the helper.
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
