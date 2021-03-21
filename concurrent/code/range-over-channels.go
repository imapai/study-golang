package main

import (
	"fmt"
)

// for 和 range为基本的数据结构提供了迭代的功能。
//也可以使用这个语法来遍历从通道中取得的值。
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for element := range queue {
		fmt.Println(element)
	}
}
