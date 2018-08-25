package main

import "fmt"

func main() {
	a := 10

	if a < 20 {
		fmt.Println("a is less than 20")
	}else{
		fmt.Println("a is bigger than 20")
	}

	switch a {
		case 10:
			fmt.Println(10)
		case 2:
			fmt.Println(2)
		default:
			fmt.Println("unknown")
	}

	//type switch

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T",i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型" )
	default:
		fmt.Printf("未知型")
	}

	fmt.Println()


	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}
