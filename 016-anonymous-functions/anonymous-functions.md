In Go its possible to declare anonymous functions and assing to a variable
```go
test := func() {
		fmt.Println("Hello")
}
test()
// output: Hello
```
anonymous functions are functions that doesn't have a name, these functions are assing to variables and then the variable can call it with the ()

---
another thing that it's possible with functions is execute the function at the end of declaration

```go
test := func(x int) int {
		return x * -1
}(16)
fmt.Println(test)
// output: -16
```

in this case we created an anonymous function that will return a number multiply by -1, and after the declaration we are doing this: **(16)**, with this we are executing the function there and the variable will recieve the returned value

---

we can pass functions as params for another function:

```go
func functionWithAFunctionAsParam(paramFunc func(int) int) {
	fmt.Println(paramFunc(7))
}
multiplyByNegativeOne := func(x int) int {
		return x * -1
}
functionWithAFunctionAsParam(multiplyByNegativeOne)
// output: -7
```

in this case the `functionWithAFunctionAsParam` recieve a params with the type `func(int) int`, this type means that it will recieve a function with an int param that returns an int, this function will call the paramFunc with the value **7** and the **print the result**, after that we created another function (this time anonymous) that multiply a number by -1, then we call the `functionWithAFunctionAsParam` function with the `multiplyByNegativeOne` function, and the result of it all is -7

---

it's possible to a function returns another function

```go
func functionThatReturnsFunction() func() {
	return func() { fmt.Println("Log from the returned function") }
}
result := functionThatReturnsFunction()
result()
// Log from the returned function
```

in this case we have the `functionThatReturnsFunction` function that at the declaration return a function without params or return values, in the function block of code we just return an anonymous function printing a message, with that we call the `functionThatReturnsFunction` function and store to a variable called **result**, this variable has the type **function**, and then we call `result()` and the inside message will print