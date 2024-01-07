To create a function in Go we use the *`func`* keyword

```go
func test(){
	fmt.Println("test")
}

test()
```

to declare a function we use `func nameOfTheFunction(){ /*function content*/ }`, and to call it you just need to do this: `nameOfTheFunction()`

---
to create a function with params we add in the () of the function

```go
func add(num1, num2 int, num3 int){
	fmt.Println(num1 + num2 + num3)
}
add(3, 8, 2)
// output 13
```

in this case we created a function called `add` with two params, **num1**, **num2** and **num3**, both of them are integers, in the call we need to pass the values of the function, in this case we passed 3, 8 and 2, which result in 13

---

to return data from a function we need to declare what type of data the function returns
```go
func add(num1, num2 int) int {
	return num1 + num2
}
result := add(3, 8)
fmt.Print(result)
// output 11
```

in this case we declare that the `add` function return an integer value passing the type after the (), and we add `int` there, so with this the function **must** return a value, if not it will not allow to execute the code, and to return the data from a function we use the keyword `return`

to return multiple data from a function we need to separate the return types between commas
```go
func addAndMultiply(num1, num2 int) (int, int) {
	return num1 + num2, num1 * num2
}
sum, multiply := add(3, 8)
fmt.Printf("sum %d, multiply: %d", sum, multiply)
// output sum 11, multiply: 24
```

in this case we declare that the function returns 2 integers values, so at the `return` keyword we must add the first value comma second value

---

it is possible to label the return values of the function, and this allow the function to work in a different way
```go
func subtractAndDivide(num1, num2 int) (subtractResponse int, divideResponse int) {
	subtractResponse = num1 - num2
	divideResponse = num1 / num2
	return
}
	subtract, division := subtractAndDivide(16, 4)
	fmt.Printf("subtract %d, division: %d", subtract, division)
  // output: subtract 12, division: 4
```
with this way we have the return as declared variables and we can assing them throughout the function, and at the `return` keyword we don't need to pass anything since we already know what variables will be returned

---

another thing that we have in functions is the `defer` functionality
```go
func functionToShowDefer(){
	defer fmt.Println("after return")
	fmt.Println("before return")
}
functionToShowDefer()
// output: 
//    before return
//    after return
```

the `defer` keyword is used inside a function to declare what will be running when the functions end, it can be used to close a connection with a database for example, in this case we used to print in what order it runs, and we can see through the output that the print after the `defer` keyword is printed earlier than the print in the `defer` keyword