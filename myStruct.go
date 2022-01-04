package main

import "fmt"

type myStruct struct {
	name string
	age  int
}

func main() {
	var a myStruct
	a.age = 18
	a.name = "glq"
	fmt.Println(a)
}
