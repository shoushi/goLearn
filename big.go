package main

import (
	"fmt"
	"math/big"
)

/*
大的整型数字是通过 big.NewInt(n) 来构造的，其中 n 位 int64 类型整数。
而大有理数是用过 big.NewRat(N,D) 方法构造。N（分子）和 D（分母）都是 int64 型整数
*/
func main() {
	a := big.NewInt(8888)
	b := a.Add(a, big.NewInt(13))
	fmt.Println(b)
	d := big.NewRat(1, 2)
	c := d.Mul(big.NewRat(1, 2), big.NewRat(3, 111))
	fmt.Println(c)
}
