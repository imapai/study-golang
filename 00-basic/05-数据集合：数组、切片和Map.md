## 数据集合
数组是一组固定大小不可变，切片slice是动态数组，map是key-value映射。

### 数组
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，
这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。

len() 和 cap() 返回结果始终一样。
#### 声明数组
- 第一种：
  - var 数组名称 [数组大小] 数组类型

- 第二种：不指定大小
  - 数组名称 := [...] 数组类型 {值1, 值2, 值3}

- 第三种：指定索引位置值，未指定为默认空值;未指定大小会根据最后索引位置初始化大小
  - arr1 := [...] string { 2:"b", 4:"d" }
  - arr2 := [6] string { 2:"b", 4:"d" }
  
测试代码：

```go
package main

import "fmt"

func main() {
	var arr [3] string
	fmt.Println(arr)
	arr2 := [...] string {"a","b"}
	fmt.Println(arr2)
	arr3 := [...] string {2:"b", 4:"d"}
	fmt.Println(arr3)
	arr4 := [6] string {2:"b", 4:"d"}
	fmt.Println(arr4)
}
```

### 切片slice
切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。

len() 和 cap() 返回结果可相同和不同。

得到切片的两种方式：

- 第一种：数组转化
- 第二种：make函数
- 第三种：字面量

测试代码：
```go
package main

import "fmt"

func main() {
	arr := [...] string {"a","b","c","d"}
	slice := arr[1:3]
	fmt.Println(slice)
	slice2 := make([] int16,4)
	fmt.Println(slice2)
	slice3 := [] string {"a","b","c"}
	fmt.Println(slice3)
}
```

#### 追加appen
```go
package main

import "fmt"

func main() {
	slice := [] string { "a", "b" }
	fmt.Println(slice)
	slice2 := append(slice,"c")
	fmt.Println(slice2)
}
```

### Map
map是key-value键值对集合

map中多有的key类型一致，所有的value类型一致，但key和value可以不一致

#### Map声明

- 使用make函数
    - 名称 := make(map[key类型] value类型)

测试代码：
```go
package main

import "fmt"

func main() {
	m := make(map[string] int16)
	m ["a"] = 1
	m ["b"] = 2
	
	for key := range m{
		fmt.Println(key, m [key])
	}
}
```    
### 常见操作
TODO