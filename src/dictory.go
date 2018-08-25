package main

import (
	"fmt"
	"os"
)

//扫描制定目录下的所有文件
func main() {

	dir, err := os.Open(".")

	if err != nil {
		//todo 处理文件夹打不开的问题
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) //-1 表示扫描改文件夹下的所有文件

	if err != nil {
		//todo 处理错误
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		fmt.Println(fileInfo.Name())
	}

}
