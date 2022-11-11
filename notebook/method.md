# method
在go中，method其实是一种函数，与函数的不同点在于多了一个receiver。

receiver是一个数据类型，可以是结构体，也可以是map，slice等。

通过下面的表达式定义
```go
func (r ReceiverType) funcName(parameters) (results)
```
上面ReceiverType是一种数据类型，method的funcName可以重复，go会根据不同的ReceiverType调用ReceiverType的func。

举个例子就清楚了，比方说有长方形，圆形等等，他们都有面积这个属性，我们可以通过函数来实现。

```go
type Circle struct {
radius float64
}

type Rectangle struct {
width, height float64
color         string
}

func (c Circle) area() float64 {
return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
return r.width * r.height
}
```

## method继承
结构体可以使用匿名参数初始化，同样的，method也是可以继承的。

比方说下面的例子，Human才有sayHi method,但由于Student也继承了Human，Human的方法
Student也可以调用，这就是method的继承。
```go
type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	object []string
}
func (h Human) sayHi() {
fmt.Println("Hi, I'm", h.name)
}
func main(){
s1 := Student{Human{"hsgao", 10}, []string{"go", "python"}}
s1.sayHi()
}
```
## method改写
上面的例子中，我们改写Student的sayHi，这就是method的改写。