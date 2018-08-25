package main

import "fmt"

func main() {

	//注释的代码会出现一个运行错误，这是因为我们先要初始化map后才能使用它
	//var x map[string]int
	//x["key"] = 10
	//fmt.Println(x)
	//
	//这里我们使用关键字make方法（和slice一样）初始化map
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	//如果我们想要查找map中的元素是否存在，我们可以直接查找对应的key值，如果没有对应的key值，则会返回0对应int，空对应string
	//这里我们可以采取一个更好的方式，使用两个变量来承接查找的key值
	//在这个例子中，第二个参数ok如果为true，则执行if condition里面的line
	if name, ok := x["nothing"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("nothing inside")
	}

	//制作一个多维map，第一个map中key是string，value为map；在第二个map中，key为string，value为string
	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}
