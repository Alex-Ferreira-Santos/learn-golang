package main

import "fmt"

func main() {
	fmt.Printf("Hello %%T: %T, %%v: %v", 10, 10)
	fmt.Println()
	fmt.Printf("Bool %%t: %t", true)
	fmt.Println()
	fmt.Printf("Number:  %%b: %b, %%o: %o, %%d: %d, %%x: %x", 10, 20, 30, 40)
	fmt.Println()
	fmt.Printf("Float: %%e: %e, %%f: %f, %%g: %g", 2.12391823, 3.2342, 0.32489274294)
	fmt.Println()
	value := fmt.Sprintf("Test %07d", 123)
	fmt.Print(value)

}
