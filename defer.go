package main

import "fmt"

func first()  {
	fmt.Println(1)
}

func second()  {
	fmt.Println(2)
}

func third()  {
	fmt.Println(3)
}

func main() {
	//defer 关键字可以让某些方法在其他方法加载完之后在执行
	//下面的例子中first先输出，然后是second
	//猜想：
	//defer会将要执行的函数放入栈中，等待其他函数执行完毕后，开始释放栈，所以这里现在栈中放入third方法，之后放入second，之后释放栈
	defer third()
	defer second()
	first()

}