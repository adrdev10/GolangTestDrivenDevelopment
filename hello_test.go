package main

import (
	"testing"
)

//Test driven development

func TestHello(t *testing.T) {
	got := Hello("Adrian")
	want := "Hello Adrian"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
