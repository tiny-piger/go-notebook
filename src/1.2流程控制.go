package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 4
	if isPer := x * x * x; isPer > 100 {
		fmt.Println("good number")
	} else if isPer == 100 {
		fmt.Println("NOT BAD")
	} else {
		fmt.Println("bad")
	}
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println(sum)
	// 可以省略前后的赋值，如果变量判断跟赋值语句是同一个的话
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)
	// 遍历map
	hashMap := map[string]int{"white": 0, "red": 1, "blue": 2}
	for key, val := range hashMap {
		fmt.Println("current key: ", key)
		fmt.Println("current val:", val)
	}
	fmt.Println("max val: ", max(1, 100, 3, 4, 5, 6, 7, 8))
	fmt.Println("\n指针变量示例")
	x = 100
	x1 := add1(&x)
	fmt.Println(x)
	fmt.Println(x1)
}

// 函数中的变参
// 类似于python中的可变参数
func max(args ...int) int {
	fmt.Println("type of args is:", reflect.TypeOf(args))
	fmt.Println("args:", args)
	max_val := args[0]
	for _, val := range args {
		if val > max_val {
			max_val = val
		}
	}
	return max_val
}

// 指针变量
func add1(x *int) int {
	*x += 1
	return *x
}
