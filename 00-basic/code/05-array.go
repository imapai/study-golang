package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr)
	arr2 := [...]string{"a", "b"}
	fmt.Println(arr2)
	arr3 := [...]string{2: "b", 4: "d"}
	fmt.Println(arr3)
	arr4 := [6]string{2: "b", 4: "d"}
	fmt.Println(arr4)
}
