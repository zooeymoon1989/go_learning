package main

import (
	"math"
	"fmt"
)

/**
可以理解为js中的函数的引用
 */
func main() {

	getSquareRoot := func(x float64) float64{

		return math.Sqrt(x)

	}

	fmt.Println(getSquareRoot(64))

}
