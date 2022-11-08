package main

import "fmt"

func main() {
	// 值传递例子
	x := 1
	add(x)
	fmt.Println("directly add val: ", x) //1

	// 指针传递例子
	pointerAdd(&x)
	fmt.Println("pointer add val:", x) //2

	// map操作
	colorMap := map[string]int{"red": 1, "blue": 2}
	addColor(colorMap)
	fmt.Println(colorMap) //map[blue:2 red:1 white:100]
}
func add(x int) int {
	x += 1
	return x
}

func pointerAdd(x *int) int {
	*x = *x + 1
	return *x
}

func addColor(color map[string]int) {
	color["white"] = 100
}
