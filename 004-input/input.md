to get an input from user in go we can use the following code

```go
  scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Print(scanner.Text())
```

here we are initializing the scanner variable with a new scanner, with the scanner.**Scan()** function we requet the input from the user, and then with scanner.**Text()** we show what the input was

anything coming from the scan will be a string, so parse the value from the scan to a int we use **strconv.ParseInt(scanner.Text(), 10, 64)**