package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	switch {
	case len(os.Args) <= 2:
		wg.Add(1)
		var input []string
		arg := os.Args[1]
		go Searchstdin(input, arg)
		wg.Wait()
	case len(os.Args) <= 3:
		args := os.Args[1]
		filename := os.Args[2]
		data, err := search(filename, args)
		if err != nil {
			fmt.Println("unable to read or file not available")
		} else {
			fmt.Println(data)
		}
	case len(os.Args) >= 4:
		wg.Add(1)
		args := os.Args[1]
		filename := os.Args[2]
		destfn := os.Args[4]
		go oflag(filename, args, destfn)
		wg.Wait()
	default:
		break

	}

}

//function to search string from standard input
func Searchstdin(input []string, arg string) error {
	defer wg.Done()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) > 0 {
			input = append(input, text)
		} else {
			break
		}
	}
	for i := range input {
		present := strings.Contains(input[i], arg)
		if present {
			fmt.Println(input[i])
		} else {
			break
		}
	}
	return nil
}

//function to search string from a file/folder
func search(filename, args string) ([]string, error) {

	var mapData = make(map[string]string)
	err := filepath.Walk(filename, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		} else {
			dt, _ := os.ReadFile(path)
			i := 0
			for range dt {
				mapData[path] += string(dt[i])
				i++
			}
			return nil
		}
	})
	var result []string
	for fn, res := range mapData {
		if strings.Contains(res, args) {
			result = append(result, fmt.Sprintf("%s: %s found in this file\n", fn, args))
		} else if !strings.Contains(res, args) {
			result = append(result, fmt.Sprintf("%s: %s is not available in this file \n", fn, args))
		}
	}
	return result, err
}

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
