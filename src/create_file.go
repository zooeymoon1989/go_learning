package main

import (
	"fmt"
	"log"
	"os"
)

const FILE_NAME = "test_file_for_open/test.txt"

func main() {

	file, err := os.OpenFile(FILE_NAME, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		file, err = os.Create(FILE_NAME)
		//log.Fatalf("failing opening file : %s" , err)
	}

	defer file.Close()

	len, err := file.WriteString("hello world! I am ready to embrace the world !")

	if err != nil {
		log.Fatalf("failed writing to file : %s", err)
	}

	//获取文件文件信息
	stat, err := file.Stat()
	if err != nil {
		//todo 处理错误
		return
	}

	fmt.Printf("\nLength: %d", len)
	fmt.Printf("\nFile Length: %d", stat.Size())
	fmt.Printf("\nFile Name: %s", file.Name())

	//file , err := os.Create(FILE_NAME)

	//if err != nil {
	//	//todo
	//	return
	//}
	//
	//defer file.Close()
	//
	//file.WriteString("i am ready")

}
