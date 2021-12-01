package main

import (
	"bytes"
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	slice := list[0:3]
	for _, v := range slice {
		fmt.Println(v)
	}
	fmt.Print(slice[0])
	fmt.Printf("slice 长度是 %d", len(slice))
	fmt.Println()
	fmt.Printf("slice 最大长度是 %d \n", cap(slice))

	//make 在数组不存在时创建切片
	slice1 := make([]int, 20)
	fmt.Printf("slice1 最大容量：%d\n", cap(slice1))

	slice2 := make([]int, 20)
	copy(slice1, slice2)
	//buffer 使用
	//这种实现方式比使用 += 要更节省内存和 CPU，尤其是要串联的字符串数目特别多的时候
	var buf bytes.Buffer
	buf.WriteString("hhhh")
	buf.WriteString("ooooo")
	fmt.Print(buf.String())
}
