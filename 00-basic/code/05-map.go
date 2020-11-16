package main

import "fmt"

func main() {
	m := make(map[string]int16)
	m["a"] = 1
	m["b"] = 2

	for key := range m {
		fmt.Println(key, m[key])
	}
}
