package main

import "fmt"

func main(){
	var mp map[string]int = map[string]int{
		"apple": 1,
		"pear": 6,
		"orange": 10,
	}
	fmt.Println(mp)

	// acces an item from map
	fmt.Println(mp["apple"])

	// change an item from map
	mp["apple"] = 100
	fmt.Println(mp["apple"])

	// delete an item from map
	delete(mp, "apple")
	fmt.Println(mp["apple"])
	fmt.Println(mp)

	// checks if an elements exists
	val, ok := mp["apple"]
	fmt.Printf("val: %d, exists: %t \n", val, ok)

	// create a new map with the make function
	mp2 := make(map[string]int)
	fmt.Print(mp2)

	for index, val := range mp{
		fmt.Printf("index: %s, value: %d \n", index, val)
	}
}