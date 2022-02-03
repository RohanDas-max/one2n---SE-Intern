package main

import (
	"os"
	"testing"
)

func TestSearchstdin(t *testing.T) {
	arg := "foo"
	input := []string{
		"foo", "bar", "foobaz", "barbazfoo", "food",
	}
	g1 := "foo"
	newi := []string{
		"foo", "bar", "foobaz", "barbazfoo", "food",
	}

	ff := func(input []string, arg string) {
		got := Searchstdin(input, arg)
		var want error = nil
		if got != want {
			t.Errorf("got %v instead of %v", got, want)
		}
	}

	ff(input, arg)
	ff(newi, g1)
}

func TestSearch(t *testing.T) {
	os.Create("test.txt")
	os.WriteFile("test.txt", []byte("Hello World"), 0400)
	filename := "test.txt"
	args := "Hello"
	defer os.Remove(filename)
	got := search(filename, args)
	if got != true || false {
		t.Errorf("got %v instead of bool", got)
	}
}

func TestOflag(t *testing.T) {
	filename := "out.txt"
	data := "any thing"

	ff := func(filename string, data string) {
		got := oflag(filename, data)
		var want error = nil
		if got != nil {
			t.Errorf("got %v instead %v", got, want)
		}
		os.Remove(filename)
	}

	ff(filename, data)
}
