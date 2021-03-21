package main

import (
	"fmt"
	"time"
)

// 定时器表示在未来某一时刻的独立事件。你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。这里的定时器将等待 2 秒。
func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("timer1 expired")

	timer2 := time.NewTimer(time.Second * 1)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()
	if timer2.Stop() {
		fmt.Println("timer2 stop")
	}
}
