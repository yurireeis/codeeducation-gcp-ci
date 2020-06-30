package main

import (
    "testing"
)

func TestSoma(t *testing.T) {
	if Soma(5,5) != 10 {
		t.Fatal("failed!")
	}
}