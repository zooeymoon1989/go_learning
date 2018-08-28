package main

import "fmt"

func main() {
	//panic可以开启一个运行时的错误
	//我们可以使用内建的函数来处理运行时的panic

	//panic("PANIC")
	//str :=recover()
	//fmt.Println(str)

	//这里有个问题：因为panic中断程序运行，我们永远看不到程序recover运行，所以无法打印输出
	//针对上述问题，我们可以使用defer关键字配合匿名函数让输出打印出来
	//这样程序不会中断，会继续执行

	defer func() {
		str :=recover()
		fmt.Println(str)
	}()

	panic("panic")
}
