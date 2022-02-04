package main

import (
	"fmt"
)

func exec() {
	n, k := flags()
	res := Controller(n, k)
	fmt.Printf("%d number player is the winner\n", res)
}

func main() {
	exec()
}
