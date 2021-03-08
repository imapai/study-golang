package main

import (
	"fmt"
	_ "fmt"
)

func zerovalue(val int) {
	val = 0
}

func zerovarible(varible *int) {
	*varible = 0
}

// 指针通过引用传递值或者数据结构
func main() {
	i := 1
	fmt.Println(i)

	zerovalue(i)
	fmt.Println(i)

	zerovarible(&i)
	fmt.Println(i)

	fmt.Println(&i)

}
