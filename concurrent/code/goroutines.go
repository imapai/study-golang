package main

import (
	"fmt"
)

func say(str string) {
	for i := 0; i < 3; i++ {
		//time.Sleep(100 * time.Microsecond)
		fmt.Println(str, ":", i)
	}
}

// Go 协程 在执行上来说是轻量级的线程。
func main() {
	say("direct")
	go say("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
