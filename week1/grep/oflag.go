package main

import (
	"flag"
	"os"
)

//function to invoke -o option to write in a file specified
func oflag(filename string, arg string, dest string) error {
	defer wg.Done()
	data, err := search(filename, arg)
	if err != nil {
		return err
	} else {
		var sr string
		for _, dt := range data {
			sr += dt
		}
		s := flag.String("o", dest, "flag to store output in a file")
		flag.Parse()
		if err := os.WriteFile(*s, []byte(sr), 0400); err != nil {
			return err
		}
		return nil
	}
}
