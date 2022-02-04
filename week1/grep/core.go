package main

import (
	"fmt"
	"os"
)

func core() {
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
