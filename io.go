package main

import (
	"bytes"
	"fmt"
)

//io包包含了一系列的方法，但是大多数都是在其他包中使用
//这里有两个接口很重要Reader和Write
//Readers支持从Read方法里面读
//Writers支持从Write方法里面写
//在go中很多方法都使用Readers和Writers作为参数

func main() {

	var buf bytes.Buffer //初始化buffer
	buf.Write([]byte("test"))

	fmt.Println(buf.String())
	// -> test

}
