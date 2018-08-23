package main

import "fmt"

type Profile struct {
	name  string
	age   uint
	hobby string
}

type DoSomething struct {
	t int
}

//在结构体中声明方法，需要在方法名前面限定要声明的结构体
//这个例子中ds为DoSomething的指针，这样在方法内部可以用结构体的成员
func (ds *DoSomething) doTest() int {
	return ds.t * 15
}

func (ds *DoSomething) getPoint() int {
	return ds.t
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("hello ! My name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {

	android := new(Android)
	android.Name = "android"
	//这里因为可以想象继承关系，android继承person，所以可以调用父类的talk方法
	android.Talk()

	fmt.Println("--------------")

	ds := &DoSomething{12}

	fmt.Println(ds.doTest())

	ds.t = 14

	fmt.Println(ds.doTest())

	fmt.Println(ds.getPoint())

	fmt.Println("--------")

	Ainfomation := Profile{"he", 12, "hahah"}

	something := &Ainfomation

	fmt.Println(*something)

	var a = Profile{"haha", 25, "heheh"}

	fmt.Println(a.age)

	fmt.Println(Ainfomation.name)

}
