package hello

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

		got := Hello("Ike", "english")
		want := "Hello, Ike"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'hello world' when an empty string in supplied", func(t *testing.T) {

		got := Hello("", "spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)

	})

}
