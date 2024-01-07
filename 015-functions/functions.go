package main

import "fmt"

func addAndMultipy(num1, num2 int) (int, int) {
	return num1 + num2, num1 * num2
}

func subtractAndDivide(num1, num2 int) (subtractResponse int, divideResponse int) {
	subtractResponse = num1 - num2
	divideResponse = num1 / num2
	return
}

func functionToShowDefer(){
	defer fmt.Println("after return")
	fmt.Println("before return")
}

func main() {
	sum, multiply := addAndMultipy(3, 8)
	fmt.Printf("sum %d, multiply: %d\n", sum, multiply)

	subtract, division := subtractAndDivide(16, 4)
	fmt.Printf("subtract %d, division: %d\n", subtract, division)

	functionToShowDefer()
}
