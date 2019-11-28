package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 1)
	want := 2
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSub(t *testing.T) {
	got := Sub(2, 1)
	want := 1
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
