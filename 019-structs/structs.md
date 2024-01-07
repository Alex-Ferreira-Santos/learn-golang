Structs in Go are used to define a custom type, it can be see as a class type if you saw it in other programming language, this can be used to create a variable with a different structure, like if we would create a variable of type Person, that need a name and an age

```go
type NameOfStruct struc {
	param1 int
	param2 int
}

var varaibleWithStructType NameOfStruct = NameOfStruct{2,5}
fmt.Println(varaibleWithStructType)
// output: {2,5}
```

to access a param from a struct we can use a dot in the variable

```go
type Point struct {
	num1 int
	num2 int
}
var point Point = Point{2,5}
var point2 Point = Point{-7,10}
fmt.Printf("point1 num1: %d, point2 num1: %d\n", point.num1, point2.num1)
// output: point1 num1: 2, point2 num1: -7
``` 

we can change the a param as any other variable
```go
type Point struct {
	num1 int
	num2 int
}
var point Point = Point{2,5}
point.num1 = 100
fmt.Printf("point1 num1: %d\n", point.num1)
// output: point1 num1: 100
``` 

we can initialize a struct specifing the values for each params, we don't need to follow the order, and if a param don't receive a value it will initialize with the default type value (for int it will be 0)
```go
type Point struct {
	num1 int
	num2 int
}
var point Point = Point{num2: 3}
fmt.Printf(point)
// output: {0 3}
``` 

and with a struct it's possible to use a struct as a type of a param in another struct
```go
type Point struct {
	num1 int
	num2 int
}

type Circle struct {
	radius float64
	center *Point
}
point5 := &Point{num2: 3}
circle := Circle{radius: 4.56, center: point5}
fmt.Println(circle)
// output: {4.56 0xc000194090}
fmt.Println(circle.center)
// output: &{0 3}
``` 
