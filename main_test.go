package main

import "testing"


func TestHellos(t *testing.T) {
	t.Run("Saying Hello World", func(t *testing.T) {
		got := hello()
		want := "Hello, world"

		if got != want {
			t.Errorf("got: %q want: %q", got, want)
		}
	})	

	t.Run("Saying Hello Chirs", func(t *testing.T) {
		got := helloName("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got: %q want: %q", got, want)
		}
	})	
}