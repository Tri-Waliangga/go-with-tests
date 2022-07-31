package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Tri", "")
		want := "Hello, Tri"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Tri", "Spanish")
		want := "Hola, Tri"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Tri", "French")
		want := "Bonjour, Tri"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Tri", "Indonesia")
		want := "Hai, Tri"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleTestHello() {
	got := Hello("Tri", "Indonesia")
	fmt.Println(got)
	// Output: Hai, Tri
}
