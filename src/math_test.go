package main

import (
	"testing"
)

// call the Add function

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d, want 3", result)
	}
}
