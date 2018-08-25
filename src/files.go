package main

import (
	"fmt"
	"os"
)

//读取文件的方法
func main() {

	file, err := os.Open("test_file_for_open/i_am_waiting_for_open.txt")

	if err != nil {
		//todo 如果有错误，则处理错误
		return
	}

	//操作结束后关闭文件句柄
	defer file.Close()

	//获取文件文件信息
	stat, err := file.Stat()
	if err != nil {
		//todo 处理错误
		return
	}

	//读取文件
	//使用切片
	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)

	if err != nil {
		//todo 无法读取文件
		return
	}

	//转化成字符串
	str := string(bs)

	fmt.Println(str)

}
