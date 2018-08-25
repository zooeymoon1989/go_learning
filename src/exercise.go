package main

import (
	"fmt"
	"time"
)

func main() {

	//找出最小值
	fmt.Println(time.Now())
	x := []int{48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
		3,
	}

	smallest := x[0]

	//for循环
	//for i:=1 ; i < len(x) ; i++ {
	//	if smallest > x[i] {
	//		smallest = x[i]
	//	}
	//}

	//while循环
	i := 1
	for i < len(x) {
		if smallest > x[i] {
			smallest = x[i]
		}
		i++
	}

	fmt.Println(smallest)
	fmt.Println(time.Now())
}
