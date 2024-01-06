package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Type your age")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	if age >= 18 {
		fmt.Print("You can drive")
	} else if age >= 16 {
		fmt.Print("You can't drive, wait some years")
	} else {
		fmt.Print("You can't drive at all")
	}
}
