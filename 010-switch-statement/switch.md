in Go we can create a switch statement with the following syntax:

```go
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
```
In this case we check the option that the user typed, if the user typed 1 it will enter inside the case with value "1", if the user typed "2" it will enter case "2", if the user type a option different than 1, 2 or 3 it will enter the default case where it going to show "Invalid option"