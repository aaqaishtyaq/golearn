package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello in 'English' when no language provided", func(t *testing.T) {
		got := Hello("Tom", "")
		want := "Hello, Tom"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello in French", func(t *testing.T) {
		got := Hello("Jean", "French")
		want := "Bonjour, Jean"

		assertCorrectMessage(t, got, want)
	})
}
