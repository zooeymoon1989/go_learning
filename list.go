package main

import (
	"container/list"
	"fmt"
)

//在go中， container/list包实现了一个双向链表

func main() {

	var x list.List

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	x.PushBack(4)

	//for e := x.Front() ; e != nil ; e = e.Next() {
	//	fmt.Println(e.Value.(int))
	//}

	var e = x.Front()

	for e != nil {

		fmt.Println(e.Value.(int))
		e = e.Next()

	}

}
