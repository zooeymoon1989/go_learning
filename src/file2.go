package main

import (
	"fmt"
	"io/ioutil"
)

//使用ioutil包可以简化读文件的方法
func main() {

	bs, err := ioutil.ReadFile("test_file_for_open/i_am_waiting_for_open.txt")

	if err != nil {
		//todo 处理错误
		return
	}

	str := string(bs)

	fmt.Println(str)

}
