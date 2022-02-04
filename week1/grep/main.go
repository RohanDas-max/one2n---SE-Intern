package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	core()
}
