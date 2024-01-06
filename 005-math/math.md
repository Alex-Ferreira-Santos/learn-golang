In Go there's this math operators

**"+"** = plus 

**"-"** = minus

"*" = multiply

**"/"** = divide

"%" = mod

in Go we can only do the operations if variables are from same type

```go
  // it work
	var num1 int = 8
	var num2 int = 4
	answer:= num1 / num2
	fmt.Print(answer)


  // doesn't work

  var num1 float32 = 8
	var num2 int = 4
	answer:= num1 / num2
	fmt.Print(answer)
```

to convert the type of the variable we can use the something like this:

```go
	var num1 float32 = 8
	var num2 int = 4
	answer:= num1 / float32(num2)
	fmt.Print(answer)
```

if you store the result of a math operation in a variable, the value will follow the type of the operation, example:


```go
	var num1 int = 9
	var num2 int = 4
	answer:= num1 / num2
	fmt.Print(answer)
  // output: 2 because asnwer will have the int type, like num1 and num2
```