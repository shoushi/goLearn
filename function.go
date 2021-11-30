package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	a, b := multiRes(1, 2)
	// 空白符 _
	c, _ := multiRes(1, 2)
	fmt.Println(c)
	func() {
		_, filename, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", filename, line)
	}()
	fmt.Printf("相加:%d,相乘：%d", a, b)
	fmt.Println()
	funMethod(1, add)
	dosth()
}

func add(a int, b int) int {
	return a + b
}

func multiRes(a int, b int) (int, int) {
	return a + b, a * b
}

func funMethod(a int, b func(int, int) (res int)) {
	fmt.Println(a + b(2, 3))
}

// 匿名函数 闭包使用
// func(){}() 最后一个() 为调用
func dosth() {
	a := 12
	func() {
		fmt.Println(a)
	}()
	defer func() {
		fmt.Println(a)
	}()
	a += 12
	func() {
		fmt.Println(a)
	}()
}
