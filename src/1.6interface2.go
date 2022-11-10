package main

import "fmt"

type Cat struct{}

type Map map[int]string

type Duck interface {
	Walk()
	Run()
}

func (c *Cat) Walk() {
	fmt.Println("walking")
}

func (c *Cat) Run() {
	fmt.Println("Runging")
}

func main() {
	var d Duck = &Cat{}
	d.Walk()
	d.Run()
	// interface 参数不是任意参数

	var c *Map
	fmt.Println(c == nil)
	fmt.Println(NilOrNot(c))
}

func NilOrNot(i interface{}) bool {
	return i == nil
}
