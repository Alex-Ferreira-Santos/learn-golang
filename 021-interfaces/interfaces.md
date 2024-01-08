Interfaces in Go are used to define that a variable has a custom behavior, like others programming languages

```go
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

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main(){
	circle1 := circle{4.5}
	rectangle1 := rect{5, 7}
	shapes := []shape{circle1, rectangle1}

	fmt.Println(shapes[0].area())
  // output: 63.61725123519331
	fmt.Println(shapes[1].area())
  // output: 35
}
```

in this case we have 2 structs, `circle` and `rect`, and both of these structs have the `area()` method, in the main function we create a circle and a rectangle, then we create a variable called `shapes`, this variable recieve a type that is an interface, this means that the variable has the behavior declared in the interface, in this case the interface `shape` has only the method area, so the variables `shapes` as an array of shape can only use the area method from every item inside the array, not matter if it's a circle or a rectangle, and because of the interface has only the area() function, so we can't access witdth, height or radius