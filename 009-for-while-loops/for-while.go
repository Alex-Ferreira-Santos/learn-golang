package main

import "fmt"

func main() {
	x := 3
	for x < 5 {
		fmt.Println()
		fmt.Printf("current x value %d", x)
		x++
	}

	for y := 0; y <= 5; y++ {
		fmt.Println()
		fmt.Printf("current y value %d", y)
		x++
	}

	a := 0
	for {
		a++
		if a == 0 {
			fmt.Println("current a value is 0")
			continue
		}
		fmt.Println("Skipped continue")
		if a > 2 {
			fmt.Println("current a value is bigger than 2 the loop will break")
			break
		}
	}
}
