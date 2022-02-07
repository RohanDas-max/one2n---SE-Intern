package main

func Controller(n, k int) int {
	if n == 1 {
		return 1
	} else {
		return ((Controller(n-1, k) + (k - 1)) % n) + 1
	}
}
