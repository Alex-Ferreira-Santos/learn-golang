package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Select an option: 1, 2 or 3")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	typedOption := scanner.Text()

	switch typedOption {
	case "1":
		fmt.Println("The Selected option was 1")
	case "2":
		fmt.Println("The Selected option was 2")
	case "3":
		fmt.Println("The Selected option was 3")
	default:
		fmt.Println("Invalid option")
	}
}
