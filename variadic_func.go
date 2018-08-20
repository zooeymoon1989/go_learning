package main

import "fmt"

//使用一个参数可变的函数

func add(variadic ...int) (int) {

	total := 0

	//for循环中key为_
	for _,v := range variadic{
		total += v
	}

	return total

}

func main() {
	fmt.Println(add(4,5,6,7,8,9,10,11))
}

