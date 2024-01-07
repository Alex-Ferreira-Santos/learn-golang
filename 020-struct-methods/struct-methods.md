When we have structs we sometime will need some methods to run some operations with that struct, for example calcule the average from a student grades

```go
type Student struct {
	name   string
	grades []int
	age    int
}

func (student Student) calculateAverage() int {
	var sum int
	for _, value := range student.grades {
		sum += value
	}
	return sum / len(student.grades)
}

student:= Student{"Alex", []int{77, 83,94}, 20}
fmt.Println(student.calculateAverage())

```

in this case we created a struct `Student` with params `name, grades and age`, and then we created a method for the `Student` struct called `calculateAverage`, then we created a variable called student and print the return value of `student.calculateAverage()`.

to create a method to a struct we need the following syntax:

```go
// func (variableReferingTheStruct StructName) functionName() returnType {}
func (student Student) calculateAverage() int 
```

the thing that includes the method at the struct is place the **`(variableReferingTheStruct StructName)`** between `func` and `functionName`