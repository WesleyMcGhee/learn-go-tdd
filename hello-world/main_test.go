package main

import "testing"


func TestHellos(t *testing.T) {
	t.Run("Saying Hello World", func(t *testing.T) {
		got := helloName("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})	

	t.Run("Saying Hello Chirs", func(t *testing.T) {
		got := helloName("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})	
	
	t.Run("in spanish", func(t *testing.T) {
		got := helloName("Juan", "spanish")
		want := "Hola, Juan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := helloName("Piere", "french")
		want := "Bonjour, Piere"

		assertCorrectMessage(t, got, want)
	}) 
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}