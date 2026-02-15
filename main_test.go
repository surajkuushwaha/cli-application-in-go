package main

import "testing"

func TestCountWord(t *testing.T) {
	input := "one two three four five"
	wants := 5

	result := CountWords([]byte(input))

	if result != wants{
		t.Fail()
	}
}
