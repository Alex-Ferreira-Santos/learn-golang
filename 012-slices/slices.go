package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var slice1 []int = x[2:4]
	fmt.Printf("slice1 value %d, length: %d, capacity: %d", slice1, len(slice1), cap(slice1))
	fmt.Println()

	var slice2 []int = x[2:]
	fmt.Printf("slice2 value %d, length: %d, capacity: %d", slice2, len(slice2), cap(slice2))
	fmt.Println()

	var slice3 []int = x[:3]
	fmt.Printf("slice3 value %d, length: %d, capacity: %d", slice3, len(slice3), cap(slice3))
	fmt.Println()

	var slice4 []int = x[:]
	fmt.Printf("slice4 value %d, length: %d, capacity: %d", slice4, len(slice4), cap(slice4))
	fmt.Println()


	var slice5 []int = []int{1, 2, 3, 4, 5}
	newSlice := append(slice5, 10)
	fmt.Printf("slice5 value %d, length: %d, capacity: %d", slice5, len(slice5), cap(slice5))
	fmt.Println()
	fmt.Printf("newSlice value %d, length: %d, capacity: %d", newSlice, len(newSlice), cap(newSlice))
	fmt.Println()

	var slice6 = make([]int, 5)
	fmt.Printf("slice6 value %d, length: %d, capacity: %d", slice6, len(slice6), cap(slice6))

}
