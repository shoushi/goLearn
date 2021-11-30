package main

import (
	"fmt"
)

func main() {
	var str = [10]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(str[0])
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %d \n", pos, char)
	}
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
}
