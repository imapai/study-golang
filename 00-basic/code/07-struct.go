package main

import (
	"fmt"
	_ "fmt"
)

type person struct {
	name string
	age  int8
}

func main() {
	p := person{
		name: "zs",
		age:  13,
	}
	fmt.Println(p.name, p.age)

	sp := &p
	fmt.Println(sp.age)
	sp.age = 18
	fmt.Println(sp.age)
}
