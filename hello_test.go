package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mahen")
	want := "Hello, Mahen"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
