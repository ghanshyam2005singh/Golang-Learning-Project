package main

import "testing"

func TestBookTickets(t *testing.T) {
	remaining := 10
	got := bookTickets("Alice", 3, remaining)
	want := 7

	if got != want {
		t.Errorf("bookTickets() = %v; want %v", got, want)
	}
}
