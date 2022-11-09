package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
	color         string
}

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	object []string
}

type Week map[string]int

func main() {
	c := Circle{2.4}
	fmt.Println("area of circle is: ", c.area())
	rect := Rectangle{3, 5, "white"}
	fmt.Println("area of rect is: ", rect.area())
	week := Week{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    7,
	}
	fmt.Println(week.getDay("Thursday"))

	rect.setColorByPointer("red")
	fmt.Println(rect)
	rect.setColorByPointer("gg")
	fmt.Println(rect)

	// method 继承
	s1 := Student{Human{"hsgao", 10}, []string{"go", "python"}}
	s1.sayHi()
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (w Week) getDay(x string) int {
	return w[x]
}

func (r *Rectangle) setColorByPointer(color string) {
	r.color = color
}

func (r Rectangle) setColorByCopy(color string) {
	r.color = color
}

func (h Human) sayHi() {
	fmt.Println("Hi, I'm", h.name)
}

func (s Student) sayHi() {
	fmt.Println("Hi, I'm student", s.name)
}
