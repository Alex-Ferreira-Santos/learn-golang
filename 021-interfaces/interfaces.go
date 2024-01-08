package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}


type circle struct {
	radius float64
}

type rect struct {
	width float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main(){
	circle1 := circle{4.5}
	rectangle1 := rect{5, 7}
	shapes := []shape{&circle1, &rectangle1}

	for _, value := range shapes {
    fmt.Println(getArea(value))
  }
}