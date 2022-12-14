package main

import (
	"errors"
	"fmt"
)

var isActivate bool

const (
	x = iota
	y
	z
)

func main() {
	//布尔值
	flag := true
	fmt.Println(isActivate)
	fmt.Println(flag)
	//	数值类型
	var a int = 100
	//var b int32 = 1000
	fmt.Println(a)
	//fmt.Println(a + b)
	// 字符串
	s := "string"
	s = "data"
	s = s + " world"
	fmt.Println(s)
	c := []byte(s)
	c[0] = 'g'
	s2 := string(c)
	fmt.Println(s2)
	// 错误类型
	err := errors.New("not suported type!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("x= ", x)
	fmt.Println("y= ", y)
	fmt.Println("z= ", z)
	arr := [5]string{"1", "2", "3", "4", "5"}
	arr[2] = "100"
	fmt.Println(arr)
	arr2 := [2][2]int{{1, 2}, {2, 3}}
	fmt.Println(arr2)
	sli1 := []int{1, 2, 3}
	fmt.Println(len(sli1))
	fmt.Println(cap(sli1))
	sli1 = append(sli1, 10)
	fmt.Println(len(sli1))
	fmt.Println(cap(sli1))
	// map
	d := map[string]int{"white": 0, "black": 1, "red": 2, "blue": 3, "green": 4, "purple": 5}
	fmt.Println(d)
	fmt.Println(d)
	fmt.Println(d)
	delete(d, "white")
	fmt.Println(d)
	delete(d, "sfd")
	val, ok := d["green"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println(val)
	}
}
