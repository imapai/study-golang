package main

import "fmt"

func main() {
	slice := []string{"a", "b"}
	fmt.Println(slice)
	slice2 := append(slice, "c")
	fmt.Println(slice2)
}
