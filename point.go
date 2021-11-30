package main

import "fmt"

func main() {
	a := 1
	c := 1
	fmt.Println(&a)
	fmt.Println(&c)
	fmt.Println(&a == &c)
	b := &a
	fmt.Println(*b)
}
