package main

import "flag"

func flags() (int, int) {
	n := flag.Int("n", 0, "option to get the numbers players")
	k := flag.Int("k", 0, "option to get the skips")
	flag.Parse()
	return *n, *k
}
