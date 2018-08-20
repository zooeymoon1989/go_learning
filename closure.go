package main

import "fmt"

func main() {

	//建立一个普通的闭包
	//x := 0
	//
	//increment := func() int {
	//	x ++
	//	return x
	//}
	//
	//fmt.Println(increment())
	//fmt.Println(increment())

	nextEven :=makeEvenGenerate()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

/**
 *闭包函数，函数名称后添加一个匿名函数,用来调用使用
 *变量赋值闭包函数，之后调用本身
 */
func makeEvenGenerate() func() (uint)  {

	i := uint(0)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}

}