package main

import "fmt"

func main() {

	fmt.Println(sumNumber(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

func sumNumber(args ...int) int {

	sum := 0

	//range关键字用于数组迭代，在for循环中，第一个参数为key，第二个参数为value，这个例子用来计算求和，所以用不到key，所以我们使用空白符"_"省略了。
	for _, num := range args {
		sum += num
	}

	return sum

}
