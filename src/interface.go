package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1, x2, y1, y2 float64
}

type Shape interface {
	area() float64
	distance() float64
}

type Circle struct {
	r float64
}

func (c *Circle) distance() float64 {
	return math.Pi * c.r * 2
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) distance() float64 {

	a := r.x2 - r.x1
	b := r.y2 - r.y1

	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

}

func (s *Rectangle) area() float64 {
	l := s.x2 - s.x1
	w := s.y2 - s.y1

	return l * w
}

//这个例子中，我们写一个方法来计算图形所有面积和

//func totalArea(circle ... Circle) (total float64) {
//	for _ , c := range circle {
//		total += c.area()
//	}
//
//	return
//}

//这里我们不能在变量可变的函数中添加两种类型的变量

//func totalArea(circle ... Circle , rectangle ... Rectangle) (total float64) {
//	for _ , c := range circle {
//		total += c.area()
//	}
//
//	return
//}

//所以我们在这里要修改方法，让输入参数支持这两种图形
//func totalArea(circles []Circle , shapes []Shape) (total float64) {
//	for _ , c := range circles {
//		total += c.area()
//	}
//
//	for _ , s := range shapes {
//		total += s.area()
//	}
//
//	return
//}

//但是这里还有一个问题，如果我们要添加新的图形，那么这个函数参数还是需要变化，需要重写
//我们知道这两个图形都有area方法，它们都实现了Shape接口

func totalArea(shapes ...Shape) (total float64) {
	for _, s := range shapes {
		total += s.area()
	}

	return
}

func totalDistance(shapes ...Shape) (total float64) {
	for _, d := range shapes {
		total += d.distance()
	}

	return
}

//接下来我们还可以优化代码，想象一下结构体Shape可能有多种图形，我们可以使用一个更大的抽象Shape结构体来包装这一层，同时在结构体中使用数组来保存这些Shape继承的图形
//这里我们创建了一个MultiShape
//同时创建了一个方法MultiTotalArea
//因为MultiTotalArea方法限定在MultiShape ，所以可以理解为类中的方法
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) multiTotalArea() (total float64) {
	for _, s := range m.shapes {
		total += s.area()
	}

	return
}

func main() {

	r := Rectangle{1, 2, 3, 5}

	fmt.Println(r.area())

	fmt.Println(r.distance())

	c := Circle{12}

	fmt.Println(c.area())

	//以下的函数是用来计算图形的总面积和
	fmt.Println(totalArea(&r, &c))

	fmt.Println(totalDistance(&r, &c))

	//这里我们创建一个MultiShape结构体，并且对它进行初始化，之后在调用其方法
	multiShape := MultiShape{
		shapes: []Shape{
			&r,
			&c,
		},
	}

	fmt.Println(multiShape.multiTotalArea())

}
