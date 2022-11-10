# interface
从定义来看，interface是一组method的组合。通过interface来定义对象的行为。
interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

## 空interface

## case demo

先粗略填写，后面补充。

## interface的数据结构
go根据interface包不包含一组method分为两种类型
+ 使用 runtime.iface 结构体表示包含方法的接口
+ 使用 runtime.eface 结构体表示不包含任何方法的 interface{} 类型；

runtime.eface 在go中是这样定义的。
```go
type eface struct { //
	_type *_type
	data  unsafe.Pointer
}
```
由两个部分组成，一个是指向底层数据的指针，一个是指向类型的指针。