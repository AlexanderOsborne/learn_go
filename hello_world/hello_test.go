package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
			t.Helper()
			if got != want {
					t.Errorf("got %q want %q", got, want)
			}
	}

	t.Run("saying hello to people", func(t *testing.T) {
			got := Hello("Chris","")
			want := "Hello, Chris"
			assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
			got := Hello("","")
			want := "Hello, World"
			assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Diana", "Spanish")
		want := "Hola, Diana"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Maria", "French")
		want := "Bonjour Maria"
		assertCorrectMessage(t, got, want)
	})

}

