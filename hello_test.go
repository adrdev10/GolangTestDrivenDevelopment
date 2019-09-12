package main

import (
	"testing"
)

//Test driven development

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Adrian")
		want := "Hello Adrian"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}
