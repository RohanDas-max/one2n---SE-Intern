package main

import "flag"

func flags() (int, int) {
	var n *int
	var k *int
	n = flag.Int("n", 14, "option to get the numbers players")
	k = flag.Int("k", 2, "option to get the skips")
	flag.Parse()
	return *n, *k
}
