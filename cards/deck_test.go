package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(d))
	}

	if d[0] != "♠ A" {
		t.Errorf("Expected first card of '♠ A', but got '%v'", d[0])
	}

	if d[len(d) - 1] != "♣ K" {
		t.Errorf("Expected first card of '♣ K', but got '%v'", d[len(d) - 1])
	}
}
