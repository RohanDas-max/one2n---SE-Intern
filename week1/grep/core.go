package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

func core() error {
	switch {
	case len(os.Args) <= 2:
		arg := os.Args[1]
		err := Searchstdin(arg)
		if err != nil {
			return err
		}
	case len(os.Args) <= 3:
		args := os.Args[1]
		filename := os.Args[2]
		wg.Add(1)
		dataChan := make(chan []string)
		errChan := make(chan error)
		go func() {
			data, err := search(filename, args)
			dataChan <- data
			errChan <- err
		}()
		data := <-dataChan
		err := <-errChan
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
		err := oflag(filename, args, destfn)
		if err != nil {
			log.Fatal(err)
		}
		wg.Wait()
	default:
		break
	}
	return nil
}
