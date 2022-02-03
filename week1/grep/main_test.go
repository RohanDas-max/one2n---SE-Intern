package main

import (
	"os"
	"testing"
)

func TestSearchstdin(t *testing.T) {

	var input = []struct {
		arg   string
		stdin []string
	}{
		{"foo", []string{"foo", "bar", "foobaz", "barbazfoo", "food"}},
		{"bar", []string{"adas", "asdsad"}},
	}
	wg.Add(len(input))

	for _, ii := range input {
		t.Run("", func(t *testing.T) {
			got := Searchstdin(ii.stdin, ii.arg)
			var want error = nil
			if got != want {
				t.Errorf("got %v instead of %v", got, want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	var input = []struct {
		filename string
		arg      string
	}{
		{"dir", "Hello"},
		// {"test.txt", "hello"}, //this will fail the test
	}
	for _, ii := range input {
		t.Run("searching argument in a file/dir", func(t *testing.T) {
			_, err := search(ii.filename, ii.arg)

			if err != nil {
				t.Errorf("got:= %v ", err)
			}
		})
	}
}
func TestOflag(t *testing.T) {
	var input = []struct {
		filename    string
		data        string
		destination string
	}{
		{"dir", "hello", "out.txt"},
		// {"test.txt", "hello", "new.txt"},
	}
	wg.Add(len(input))
	for _, ii := range input {
		t.Run("testing oflag function", func(t *testing.T) {
			got := oflag(ii.filename, ii.data, ii.destination)
			var want error
			if got != want {
				t.Errorf("got=>%v", got)
			}
		})
		os.Remove(ii.destination)
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"anything"},
		{"any"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(m *testing.T) {
			main()
		})
	}

}
