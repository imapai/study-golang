package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(1000)
	fmt.Println("done")
	done <- true
}

// 我们可以使用通道来同步 Go 协程间的执行状态。这里是一个使用阻塞的接受方式来等待一个 Go 协程的运行结束。
func main() {
	done := make(chan bool, 1)
	worker(done)
	fmt.Println(<-done)
}
