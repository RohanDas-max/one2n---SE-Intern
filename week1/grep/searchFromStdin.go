package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

//function to search string from standard input
func Searchstdin(arg string) error {
	defer wg.Done()
	var input []string
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
			return errors.New("ooops!! not found")
		}
	}
	return nil
}
