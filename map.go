package main

import "fmt"

func main() {
	map1 := map[int]string{1: "asd", 2: "asda"}
	map1[3] = "i am three"
	fmt.Println(map1)
	//判断是否存在key=5
	v, a := map1[5]
	fmt.Println(v)
	fmt.Println(a)
	delete(map1, 2)
	fmt.Println(map1)

	for k, v := range map1 {
		fmt.Printf("key:%d,value:%s", k, v)
	}
}
