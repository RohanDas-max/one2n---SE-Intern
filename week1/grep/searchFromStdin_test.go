package main

import (
	"testing"
)

func TestSearchstdin(t *testing.T) {
	wg.Add(1)
	err := Searchstdin("hello")
	if err != nil {
		t.Errorf("Searchstdin() error = %v", err)
	}
}
