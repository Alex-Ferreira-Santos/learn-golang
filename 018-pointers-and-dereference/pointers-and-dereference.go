package main

import "fmt"

func changeValue(str string) {
	str = "bla"
	fmt.Printf("Inside changeValue function: %s\n", str)
}

func changeValueForReal(str *string) {
	*str = "bla"
	fmt.Printf("Inside changeValueForReal function: %s\n", *str)
}

func main() {
	num1 := 7
	fmt.Println(&num1)

	num2 := &num1
	fmt.Println(num1, num2)

	*num2 = 8
	fmt.Println(num1, num2)
	fmt.Println("==============================================")

	originalString := "originalString"

	fmt.Printf("value before changeValue(): %s\n", originalString)
	changeValue(originalString)
	fmt.Printf("value after changeValue(): %s\n", originalString)
	
	fmt.Println("==============================================")
	fmt.Printf("value before changeValueForReal(): %s\n", originalString)
	changeValueForReal(&originalString)
	fmt.Printf("value after changeValueForReal(): %s\n", originalString)

}
