package main

import "fmt"

func main() {

	//a := 20
	//var ip  *int
	//
	//ip = &a
	//
	//fmt.Println("a变量的地址是：",&a)
	//
	//fmt.Println("ip变量储存的指针地址：", ip)
	//
	//fmt.Println("*ip 变量的值：" , *ip)
	const MAX int = 3
	a := []int{10,100,200}
	var i int
	var ptr [MAX]*int;

	for  i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for  i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
	}

	x := 5
	changeValue(&x)

	fmt.Println(x)

}
//通过指针改变值
func changeValue( x *int )  {
	*x = 0
}
