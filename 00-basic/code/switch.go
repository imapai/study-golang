package main

import "fmt"

func main() {
	for num := 1; num < 7; num = num + 1 {
		fmt.Printf("当 num = %d 时，", num)
		switch_case(num)
	}
}
func switch_case(num int) {
	switch num {
	case 1:
		fmt.Printf("输出 num = 1 \n")
	case 2:
		fmt.Printf("输出num = 2 \n")
	case 3, 4, 5:
		fmt.Printf("其他 \n")
	default:
		fmt.Printf("未匹配case，默认 \n")
	}
}
