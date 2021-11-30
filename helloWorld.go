package main

import (
	"fmt"
	"time"
)

func main() {
	var a string = "hello golanguaeg"
	fmt.Println(a)
	var b, c = 1, "asd"
	fmt.Println(b, c)
	d := true
	fmt.Println(d)
	const e = "qwer"
	fmt.Printf(e)
	t := time.Now().UTC()
	fmt.Println(t)
	fmt.Println(&t)
}
