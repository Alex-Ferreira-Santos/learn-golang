package main

import "fmt"

func main(){
	var arr []int = []int{1,2,34,2,123,4532,65}
	for index, element := range arr {
		fmt.Printf("index: %d, element: %d\n", index, element)
	}
}