package main

import "fmt"

var fibs [41]uint64

func main() {
	fibonacci(40)
	for n := range fibs {
		fmt.Println(n)
	}
}

func fibonacci(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	}
	res = fibonacci(n-1) + fibonacci(n-2)
	fibs[n] = res
	return
}
