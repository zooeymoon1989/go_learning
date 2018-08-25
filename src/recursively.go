package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//递归查询所有文件
func main() {

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}
