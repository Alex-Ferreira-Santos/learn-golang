package main

import "fmt"

type Point struct {
	num1 int
	num2 int
}

type Circle struct {
	radius float64
	center *Point
}


type Circle2 struct {
	radius float64
	*Point
}

func changeNum2(pt *Point){
	pt.num2 = 100
}

func main(){
	var point Point = Point{2,5}
	point2 := Point{-7,10}
	fmt.Println(point, point2)
	fmt.Printf("point1 num1: %d, point2 num1: %d\n", point.num1, point2.num1)

	point.num1 = 100
	fmt.Printf("point1 num1: %d, point2 num1: %d\n", point.num1, point2.num1)

	point3 := Point{num2: 3}
	fmt.Println(point3)

	point4 := &Point{num2: 8}
	fmt.Printf("point4 before changeNum2: %d\n",point4)
	changeNum2(point4)
	fmt.Printf("point4 after changeNum2: %d\n",point4)

	point5 := &Point{num2: 3}
	circle := Circle{radius: 4.56, center: point5}
	fmt.Println(circle)
	fmt.Println(circle.center)

	circle2 := Circle2{radius: 4.56}
	fmt.Println(circle2.num2)
}