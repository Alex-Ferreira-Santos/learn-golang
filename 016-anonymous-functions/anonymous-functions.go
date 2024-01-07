package main

import "fmt"

func functionWithAFunctionAsParam(paramFunc func(int) int) {
	fmt.Println(paramFunc(7))
}

func functionThatReturnsFunction() func() {
	return func() { fmt.Println("Log from the returned function") }
}

func main() {
	anonymousFunc := func() { fmt.Println("Calling anonymous function") }
	anonymousFunc()
	test := func(x int) int {
		return x * -1
	}(16)
	fmt.Println(test)

	multiplyByNegativeOne := func(x int) int {
		return x * -1
	}
	functionWithAFunctionAsParam(multiplyByNegativeOne)

	result := functionThatReturnsFunction()
	result()

	functionThatReturnsFunction()()
}
