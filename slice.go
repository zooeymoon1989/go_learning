package main

import "fmt"

//slice通常用于展示列表，你可以快速的查找有序列表，比如说第十名得票
func main() {
	slice1 := []int{1, 2, 3}
	//使用make方法创建slice
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := append(slice3, 4, 5)
	fmt.Println(slice3, slice4)

	slice5 := []int{1, 2, 3, 4, 5, 6, 7}
	slice6 := make([]int, 12)
	//如果slice的长度大于，多出的部分用默认值补充
	copy(slice6, slice5)
	fmt.Println(slice5, slice6)
}
