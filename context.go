package main

import (
	"fmt"
	"time"
)

func main() {
	go threadFunc()
	go threadFunc()

}

func threadFunc() {
	for i := 0; i < 10; i++ {
		time.Sleep(2000)
		fmt.Println(i)
	}
}
