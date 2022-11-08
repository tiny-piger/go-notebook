package main

import (
	"fmt"
)

type dealInt func(int) bool

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

	// defer
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	//file, err := os.OpenFile("a.txt", os.O_CREATE, 0x666)
	//defer file.Close()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//函数作为返回值

	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(filter(sli, isOdd))
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

func isOdd(num int) bool {
	if num%2 != 0 {
		return true
	}
	return false
}

func filter(sli []int, f dealInt) []int {
	var result []int
	for _, val := range sli {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}