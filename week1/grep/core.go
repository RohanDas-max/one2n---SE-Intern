package main

import (
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func core() error {
	switch {
	case len(os.Args) <= 2:
		wg.Add(1)
		arg := os.Args[1]
		go Searchstdin(arg)
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
		args := os.Args[1]
		filename := os.Args[2]
		destfn := os.Args[4]

		wg.Add(2)
		errChan := make(chan error)
		go func() {
			err := oflag(filename, args, destfn)
			errChan <- err
		}()
		err := <-errChan
		close(errChan)
		if err != nil {
			return err
		}
		wg.Wait()
	default:
		break
	}
	return nil
}
