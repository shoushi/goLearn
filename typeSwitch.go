package main

import (
	"fmt"
)

func main() {
	var x interface{}
	switch x.(type) {
	case int:
		fmt.Print("is int")
	case nil:
		fmt.Print("is interface")
	}
}
