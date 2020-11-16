package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d"}
	slice := arr[1:3]
	fmt.Println(slice)
	slice2 := make([]int16, 4)
	fmt.Println(slice2)
	slice3 := []string{"a", "b", "c"}
	fmt.Println(slice3)
}
