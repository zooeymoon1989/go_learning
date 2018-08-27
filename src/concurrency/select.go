package main

import (
	"time"
	"fmt"
)

func main() {

	//make一个channel可以添加第二个参数，capacity为1。
	//一般来说，channel是同步的，两端都要等到另一端准备好。
	//buffer channel是异步的，接收和发送一个信息不需要等待，除非当前channel已满。
	//如果当前channel满了，放到就会等待，直到至少有一个room
	// int值为整数
	// c:= make(chan string , 2)
	// messages <- "no.1"
	// messages <- "no.2"
	// fmt.Println(messages)
	// fmt.Println(messages)
	// -----------------------
	// output:
	// 		no.1
	// 		no.2


	c1 := make(chan string)
	c2 := make(chan string)

	//每两秒输出from 1
	go func() {
		for {
			c1<- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	//每三秒输出from 2
	go func() {
		for {
			c2<- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()


	go func() {

		for {
			//select 会随机选一个case进行调用
			select {
				case msg1 := <- c1 :
					fmt.Println(msg1)
				case msg2 := <- c2 :
					fmt.Println(msg2)
				//time.after创建一个channel，在给定的期间之后，会创建一个当前时间
				case msg3 := <- time.After(time.Second):
					fmt.Println(msg3)
				//如果没有channel准备好，default马上就会调用
				//default:
				//	fmt.Println("nothing ready")
			}
		}

	}()


	var input string

	fmt.Scanln(&input)

}
