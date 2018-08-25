package main

import (
	"net"
	"fmt"
	"encoding/gob"
)

//我们可以创建一个TCP服务使用net包中的Listen函数
 //Listen使用了network类型，绑定了端口和地址，返回一个netListener
 /**
 	type Listener interface {
		// Accept waits for and returns the next connection to the listener.
		Accept() (c Conn, err error)
		// Close closes the listener.
		// Any blocked Accept operations will be unblocked and return errors.
		Close() error
		// Addr returns the listener's network address.
		Addr() Addr
	}
  */

//一旦我们拥有一个Listener，我们可以调用Accept，这个方法等待客户端连接和返回一个net.Conn
//net.Conn实现的是io.Reader和io.Writer接口，这样我们可以像读写文件一样进行操作

func server()  {
	//监听端口
	ln , err := net.Listen("tcp",":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	//while循环
	for {
		//接收一个连接
		c ,err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleServerConnection(c)
	}

}

func handleServerConnection(c net.Conn)  {
	//接收信息
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Received : ",msg)
	}

	c.Close()
}

func client()  {
	//连接到服务器
	c , err := net.Dial("tcp","127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 发送消息
	msg := "Hello, World"
	fmt.Println("Sending:", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
