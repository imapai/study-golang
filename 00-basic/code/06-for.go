package main

import "fmt"

func main() {
	person := [...]string{"apple", "banana"}
	fmt.Printf("len = %d,cap = %d,array = %v \n", len(person), cap(person), person)
	for index, value := range person {
		fmt.Printf("person[%d] -> %s \n", index, value)
	}
}
