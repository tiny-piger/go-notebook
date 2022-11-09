# method
在go中，method其实是一种函数，与函数的不同点在于多了一个receiver。

通过下面的表达式定义
```go
func (r ReceiverType) funcName(parameters) (results)
```
上面ReceiverType是一种数据类型，method的funcName可以重复，go会根据不同的ReceiverType调用ReceiverType的func。

举个例子就清楚了，比方说有长方形，圆形等等，他们都有面积这个属性，我们可以通过函数来实现。


## method继承

## method改写
