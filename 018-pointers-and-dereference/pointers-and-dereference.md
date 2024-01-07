When we declare a variable we know where that value is, inside the variable, but we don't know where the pc store the variable in memory, to see the memory location of a variable you can do this:

```go
num1 := 7
fmt.Println(&num1)
// 0xc0000120f0
```

in this case we declare the variable **num1** and then we print the variable with a **`&`** in front of it, this means that we are not using the varaible value, we are using the variable memory value, and running the code the terminal shows **0xc0000120f0**, this is where the varaible is storage inside the pc, this value can change any time we run the code, since the variable will be storage in a different memory location

---

with this we have the power to change the value from a variable through other variable, the previous folder was talking about mutable and immutable data, with this all data can be mutable

```go
num1 := 7
fmt.Println(&num1)
// output 0xc0000120f0

num2:= &num1
fmt.Println(num1, num2)
// output 7 0xc0000120f0

*num2 = 8

fmt.Println(num1, num2)
// output 8 0xc0000120f0
```

in this case we create **num1** variable and then we create **num2** variable recieving the **`pointer`** of the **num1** variable, this means that the **num2** is exacly the same as **num1** and not just a copy, then we change the value of **num2** with a different way, we do this using ***num2**, this will make a **`dereference`** to the original variable (**num1**), and changing the value in **both** of them

---

we can use **`pointers`** and **`dereference`** to change the value of a variable inside a function

```go
func changeValue(str string) {
	str = "bla"
	fmt.Printf("Inside changeValue function: %s\n", str)
}

func changeValueForReal(str *string) {
	*str = "bla"
	fmt.Printf("Inside changeValueForReal function: %s\n", *str)
}

originalString := "originalString"

fmt.Printf("value before changeValue(): %s\n", originalString)
changeValue(originalString)
fmt.Printf("value after changeValue(): %s\n", originalString)

fmt.Println("==============================================")
fmt.Printf("value before changeValueForReal(): %s\n", originalString)
changeValueForReal(&originalString)
fmt.Printf("value after changeValueForReal(): %s\n", originalString)


// output
// value before changeValue(): originalString
// Inside changeValue function: bla
// value after changeValue(): originalString
// ==============================================
// value before changeValueForReal(): originalString
// Inside changeValueForReal function: bla
// value after changeValueForReal(): bla
``` 

in this case we have 2 functions `changeValue` and `changeValueForReal`, the `changeValue` function recieve a string as param and then change the value of the string param to "bla", and the `changeValueForReal` function recieve a string pointer as param and then change the string dereference to value "bla", both of the functions do the same, change the string value, but the difference between them is that the `changeValue` will only update the string value **inside** the function, so when the function ends the original string value will continue to be the same, otherwise the `changeValueForReal` function as recieves a pointer stead of a value, it will modify the original string value, since it is changing the value in that memory location, seeing the output will help to understand