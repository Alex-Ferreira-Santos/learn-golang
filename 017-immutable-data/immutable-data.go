package main

import "fmt"

func main(){
	var num1 int = 5
	num2:= num1
	num2 = 7
	fmt.Printf("num1: %d, num2: %d \n", num1, num2)

	var slice []int = []int{3,4,5}
	newSlice := slice
	newSlice[0] = 100
	fmt.Printf("slice: %d, newSlice: %d\n",slice, newSlice)

	var mp map[string] int = map[string]int{"hello": 3}
	newMap := mp
	newMap["hello"] = 300
	fmt.Println(newMap, mp)

	var array [2]int = [2]int{3,4}
	newArr := array
	newArr[0] = 500
	fmt.Printf("array: %d, newArr: %d\n",array, newArr)

}