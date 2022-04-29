package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	// Calling the count function to the number of words
	// received from the Standard Input and printing it out
	b := bytes.NewBufferString("word1 word2 word3")

	exp := 3
	res := count(b, false)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	// Calling the count function to the number of lines
	// received from the Standard Input and printing it out
	b := bytes.NewBufferString("line1\nline2\nline3")

	exp := 3
	res := count(b, true)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}
