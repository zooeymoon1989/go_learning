package main

import "fmt"

func main() {

	//for循环
	//for i := 0 ; i < 10 ; i++ {
	//	fmt.Println(i)
	//}

	//while循环
	//i := 1
	//for i < 20 {
	//	fmt.Println(i)
	//	i ++
	//}

	number := [6]int{1,2,3,5}

	for i,x := range number{
		fmt.Println(i,x)
	}


}
