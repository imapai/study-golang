## struct 结构体

struct 结构体是各个字段类型的集合
``` go
type person struct {
	name string
	age int8
}
```
- type name struct {}创建一个名为name的结构体
- 你可以在初始化一个结构体元素时指定字段名字
- 省略的字段将被初始化为零值
- & 前缀生成一个结构体指针
- 使用点来访问结构体字段，也可对指针使用，它会自动解开
- 结构体是可变的