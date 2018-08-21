package main

import "fmt"

type infomation struct {
	name string
	age uint
	hobby string
}

func main() {

	Ainfomation := infomation{"he", 12, "hahah"}

	something := &Ainfomation

	fmt.Println(*something)

	var a = infomation{"haha",25,"heheh"}

	fmt.Println(a.age)

	fmt.Println(Ainfomation.name)

}