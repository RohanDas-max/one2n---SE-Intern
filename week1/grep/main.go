package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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
		data, _ := search(filename, args)
		fmt.Println(data)

	case len(os.Args) >= 4:
		args := os.Args[1]
		filename := os.Args[2]
		destfn := os.Args[4]
		oflag(filename, args, destfn)
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
func search(filename, args string) ([]string, error) {
	var mapData = make(map[string]string)
	if err := filepath.Walk(filename, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		dt, _ := os.ReadFile(path)
		i := 0
		for range dt {
			mapData[path] += string(dt[i])
			i++
		}
		return nil
	}); err != nil {
		return nil, err
	}
	var result []string
	for fn, res := range mapData {
		if strings.Contains(res, args) {
			result = append(result, fmt.Sprintf("%s: %s found in this file\n", fn, args))
		} else if !strings.Contains(res, args) {
			result = append(result, fmt.Sprintf("%s: %s is not available in this file \n", fn, args))
		}

	}
	return result, nil
}

//function to invoke -o option to write in a file specified
func oflag(filename string, arg string, dest string) error {

	st, err := search(filename, arg)
	var sr string
	for _, dt := range st {
		sr = dt
	}
	if err != nil {
		return err
	} else {
		s := flag.String("o", dest, "flag to store output in a file")
		flag.Parse()
		if err := os.WriteFile(*s, []byte(sr), 0400); err != nil {
			return err
		}
		return nil
	}

}
