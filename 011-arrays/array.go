package main

import "fmt"

func main() {
	fmt.Println("Creating an array")
	var arr [5]int
	fmt.Println(arr)
	fmt.Println("Accessing the first value of the array")
	fmt.Println(arr[0])
	fmt.Println("Changing the first value of the array")
	arr[0] = 100
	fmt.Println(arr)

	fmt.Println("Add default values to positions")
	arr2 := [3]int{4, 5, 6}
	fmt.Println(arr2)

	fmt.Println("Sum all elements from arr2")
	sum := 0
	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}
	fmt.Println(sum)


	fmt.Println("Creating a 2D array")

	arr3:= [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arr3)

}
