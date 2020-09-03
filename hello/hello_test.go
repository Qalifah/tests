package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("mike", empty)
		want := "Welcome, mike"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Welcome, Human' when an empty string is supplied", func(t *testing.T) {
		got := hello(empty, empty)
		want := "Welcome, Human"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := hello("ab", french)
		want := "Bonjour, ab"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := hello("jude", spanish)
		want := "Hola, jude"
		assertCorrectMessage(t, got, want)
	})
}