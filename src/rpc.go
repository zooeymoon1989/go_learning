package main

import (
	"net/rpc"
	"net"
	"fmt"
)

type Server struct {

}

//设置结构体函数
func (server *Server) Negate(i int64 , reply *int64)  error {
	*reply = -i
	return nil
}

func serverRpc()  {
	rpc.Register(new(Server))
	ln , err := net.Listen("tcp" , ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c , err := ln.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(c)
	}
}

func clientRpc() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	//rpc调用远程方法
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}
}

func main() {
	go serverRpc()
	go clientRpc()
	var input string
	fmt.Scanln(&input)
}