package main

import "fmt"

func max(num1 int , num2 int) (int) {

	if num1 > num2 {
		return num1
	}else{
		return num2
	}

}

func swap( x , y string ) (string , string)  {
	return y , x
}

func returnTwoValues(x string , y int) (int,int)  {
	return len(x) , y*10
}

func main() {

	fmt.Println(max(12,13))
	fmt.Println(swap("world" , "hello"))
	fmt.Println(returnTwoValues("hello world",15))
}