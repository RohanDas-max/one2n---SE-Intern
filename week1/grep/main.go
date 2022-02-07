package main

import "log"

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}

}
