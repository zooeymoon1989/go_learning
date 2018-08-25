package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
)

func main() {

	//创建一个hasher
	h := crc32.NewIEEE()
	//写入数据
	h.Write([]byte("hola"))
	//计算校验和
	v := h.Sum32()

	fmt.Println(v)

	//crc32 实现了一个Write接口
	//这样就可以写入byte
	//我们在完成写入之后，可以使用sum32返回一个uint32数字
	/*
	   这个方法可以用来做文件校验在下载文件之前，获取文件的uint32，下载完成之后生成一个uint32，两者做比较，这样可以验证文件的合法性
	*/



	s := sha1.New() //创建一个sha1
	s.Write([]byte("test"))
	bs := s.Sum([]byte{})
	fmt.Println(string(bs[:]))
}
