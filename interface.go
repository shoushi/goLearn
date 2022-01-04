package main

import "fmt"

type Shaper interface {
	Area() float32
}

type ib interface {
	Shaper
	say()
}
type Rectangle struct {
	length, width float32
}

func (a Rectangle) Area() float32 {
	return a.length * a.width
}

func main() {
	var sha Shaper
	sql := new(Rectangle)
	sql.length = 2.2
	sql.width = 3.3
	fmt.Print(sql.Area())
	sha = sql
	if u, ok := sha.(Rectangle); ok {
		fmt.Printf("The type of sha is: %T\n", u)
	}
}
