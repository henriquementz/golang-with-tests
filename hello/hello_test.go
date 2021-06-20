package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Ike")
	want := "Hello, Ike"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
