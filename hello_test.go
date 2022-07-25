package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tri")
	want := "Hello, Tri"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
