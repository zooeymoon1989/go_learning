package main

import (
	"fmt"
	"sort"
)

//自定义排序

type People struct {
	Name string
	Age  int
}

//实现我们自己sort方法，我们在这里可以创建一个自定义的type
//让它与我们想要排序的slice相等
type ByName []People

//todo 实现sort/sort的三个接口方法
func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []People

func (this ByAge) Len() int {
	return len(this)
}
func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {

	kids := []People{
		{"Jill", 9},
		{"Jack", 12},
		{"Zooey", 18},
		{"Mike", 14},
	}
	//如果想使用自定义sort，那么必须要实现sort接口的三个方法
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Println(kids)

}
