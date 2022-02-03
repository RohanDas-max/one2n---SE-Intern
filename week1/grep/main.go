package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	switch {
	case len(os.Args) <= 2:
		var input []string
		arg := os.Args[1]
		Searchstdin(input, arg)

	case len(os.Args) <= 3:
		args := os.Args[1]
		filename := os.Args[2]
		found := search(filename, args)
		if found {
			fmt.Printf("%s found in %s \n ", args, filename)
		}
	case len(os.Args) >= 4:
		args := os.Args[1]
		filename := os.Args[2]
		found := search(filename, args)
		if found {
			output := fmt.Sprintf("%s found in %s \n ", args, filename)
			OutputFilename := os.Args[4]
			if OutputFilename != "" {
				go oflag(OutputFilename, output)
				time.Sleep(10 * time.Millisecond)
			}
		}
	default:
		fmt.Println("Ohh Noooo!")
	}
}

//function to search string from standard input
func Searchstdin(input []string, arg string) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			input = append(input, text)
		} else {
			break
		}
	}
	for i := range input {
		present := strings.Contains(input[i], arg)
		if present {
			fmt.Println(input[i])
		}
	}
	return nil
}

//function to search string from a file/folder
func search(filename, args string) bool {
	var res string
	err := filepath.Walk(filename, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		data, _ := os.ReadFile(path)
		res = string(data)

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(res, args) {
		return true
	}
	return false
}

//function to invoke -o option to write in a file specified
func oflag(filename string, data string) error {
	s := flag.String("o", filename, "flag to store output in a file")
	flag.Parse()
	if err := os.WriteFile(*s, []byte(data), 0400); err != nil {
		return err
	}
	return nil
}
