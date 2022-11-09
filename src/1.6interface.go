package main

import (
	"fmt"
	"strconv"
)

type Human1 struct {
	name  string
	age   int
	phone string
}

type Student1 struct {
	Human1 //匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human1  //匿名字段
	company string
	money   float32
}

// Human1实现SayHi方法
func (h Human1) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human1实现Sing方法
func (h Human1) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// Employee重载Human1的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men被Human1,Student1和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func (h Human1) String() string {
	return "《" + "name:" + h.name + ",age:" + strconv.Itoa(h.age) + "》"
}

type List interface{}

func main() {
	mike := Student1{Human1{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student1{Human1{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human1{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human1{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student1
	i = mike
	fmt.Println("This is Mike, a Student1:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

	var a interface{}
	//s := 100
	a = []string{"1", "2"}
	fmt.Println(a)
	fmt.Println(mike)
	val, ok := a.([]string)
	if !ok {
		fmt.Println("error")
		fmt.Println(val)
	}
	fmt.Println(val)

	var list [3]List
	list[0] = 100
	list[1] = "string"
	list[2] = []string{"a", "b"}

	for _, val := range list {
		switch valType := val.(type) {
		case int:
			fmt.Println(val, "type is:", valType)
		case string:
			fmt.Println(val, "type is:", valType)
		case [2]string:
			fmt.Println(val, "type is:", valType)
		default:
			fmt.Println("others")
		}
	}

}
