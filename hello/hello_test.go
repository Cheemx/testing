package main

import "testing"

func TestHello(t *testing.T) {
	// Subtest to test regular functioning of Hello
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Cheems", "")
		want := "Hello, Cheems"

		assertCorrectMessage(t, got, want)
	})
	// Subtest to test empty strings in Hello
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Frank", "French")
		want := "Bonjour, Frank"

		assertCorrectMessage(t, got, want)
	})
}

// Refactored the assertion into a helper function!
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this is a helper method. By doing this, when it fails, the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems more easily.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
