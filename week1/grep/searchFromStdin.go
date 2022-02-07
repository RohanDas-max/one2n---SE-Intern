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
	err := Checking(input, arg)
	if err != nil {
		return err
	}
	return nil
}

func Checking(input []string, arg string) error {
	var present bool
	var res []string
	for i := range input {
		present = strings.Contains(input[i], arg)
		if present {
			res = append(res, input[i])
		}
	}
	if res != nil {
		fmt.Println(res)
	} else {
		return errors.New("oops! not found")
	}
	return nil
}
