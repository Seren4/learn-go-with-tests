package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		expected := "Hello, Chris"

		if got != expected {
			t.Errorf("got %q, want %q", got, expected)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		expected := "Hello, World"

		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	})

}
