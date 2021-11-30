package main

import (
	"log"
)

func main() {
	connectDb()
}

/*
	defer 关键字保证修饰的语句在return后执行
	defer是在函数结束时调用，但是defer 函数参数确是立即求值的。
*/
func connectDb() (res int) {
	// 参数直接赋值，结果 0
	defer log.Println(res)
	// return 后执行函数内的参数调用，结果 2
	defer func() {
		log.Println(res)
	}()
	// 函数参数直接赋值，结果 0
	defer func(b int) {
		log.Println(b)
	}(res)
	return 2
}
