package hello

import "testing"

func TestHello(t *testing.T)  {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		//t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("wqb", func(t *testing.T) {
		got := Hello("haha1")
		want := "Hello, haha"
		assertCorrectMessage(t,got,want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("World")
		want := "Hello, World"

		assertCorrectMessage(t,got,want)
	})

}