package main

import "fmt"

func main() {
	// 结构体及其初始化方式
	type Person struct {
		name string
		age  int
	}
	type Student struct {
		Person
		subject []string
		name    string
	}
	var p Person
	p.name = "hsgao"
	p.age = 20
	fmt.Println(p)
	p2 := Person{"bob", 20}
	p3 := Person{name: "judy", age: 10}
	fmt.Println(p2)
	fmt.Println(p3)
	s1 := Student{Person{"bob", 20}, []string{"it", "english"}, "hhhh"}
	fmt.Println(s1.name)
	fmt.Println(s1.Person.name)
}
