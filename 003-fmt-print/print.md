In go we can print a value some ways, all of them uses the fmt module

to print a string in a single line we can use:

```go
fmt.print("Hello")
// output: Hello
```
this will print "Hello" in one line, if another `fmt.print()` function is called the output will join the same line as "Hello"

to print a string with a value inside we can use the `fmt.printf()` function:

```go
name := "Alex"
fmt.printf("Hello %v", name)
// output: Hello Alex
```

with this `fmt.printf()` function we have more options to show a value in the string

**%v** will replace the string position for a variable after a comma
**%T** will show you the type of the value after a comma
**%%** will show the % symbol, if you use only one % it can generate a error in the string

**t** is used to insert a boolean value inside the string, it only accepts true or false

**%b** is used to print a number in base 2 (binary)
**%o** is used to print a number in base 6
**%d** is used to print a number in base 10
**%x** is used to print a number in base 16

**e%** is used to print a number in scientific notation
**%F** is used to print a decimal number with no exponent
**%g** is used for large exponents

**s%** is used to print a string
**%q** is used to print a double quote string

to format the string variable with the **%f** you can use number between the **%** and **f**

example:

```go
fmt.printf("%9.2f", 2.3245354)
// output:     2.32
```

if you want to show a number without precision you can use only a **.** between the **%** and **f**

```go
fmt.printf("%.f", 2.3245354)
// output: 2
```

if you want to store the output from the printf function you can use `fmt.Sprintf` and store into a variable:

```go
value := fmt.Sprintf("Test %07d", 123)
fmt.Print(value)
// output: 0000123
```
