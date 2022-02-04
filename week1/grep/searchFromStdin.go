package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
