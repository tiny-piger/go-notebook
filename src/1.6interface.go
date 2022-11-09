package main

import "fmt"

func main() {
	var a interface{}
	var i int = 5
	s := "Hello world"
	// a可以存储任意类型的数值
	a = i
	fmt.Println(a)
	a = s
	fmt.Println(a)
}
