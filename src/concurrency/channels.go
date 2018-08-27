package main

import (
	"fmt"
	"time"
)

//Channels provide a way for two goroutines to communicate with each other and synchronize their execution.

//我们可以制定channel类型的方向，下面的函数中我们规定pinger只能发送到c
func pinger(c chan<- string)  {
	for i := 0 ; ; i++ {
		// <- 表示是用来在channel中接收和发送信息

		c <- "ping"  //代表发送ping信息
	}
}

func ponger(c chan<- string)  {

	for i:= 0 ; ; i ++ {
		c <- "pong"
	}

}
//规定printer只能接受信息
func printer(c <-chan string) {
	for {
		msg := <- c  //接收信息，并且把它存到msg中
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c = make(chan string)
	//当pinger尝试发送信息的时候，如果没有接收方，它会等待知道printer接收完信息-----也就是说这里会阻塞（blocking）
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
