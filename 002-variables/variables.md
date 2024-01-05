## Declaration

---

Variables in Go are declared with the **var** word

The structure of a variable has the follow partern

```go
// var name type = value
   var greeting string = "Hello"
```

Once the variable was initialized you can reassing the value just making a simple

```go
greeting = "Hello there!"
```

A variable doesn't need to initialize with a value, you can define the name and the type and later on assing a value to it

```go
var name string
name = "Alex"
```


### **Attention!!! all default variables in Go for numbers are 0, and bool are false**

```go
	var name string
	var num uint8
	var decimalNumber float32 
	var isValid bool
	fmt.Println(name, num, decimalNumber, isValid)
    // output 0 0 false

```


## Variable Types

---

In go we have a lot of different variable types for numbers, there's a specific type for each size of number, for example the **byte/uint8** type allow us to assing a variable with a number between **0** to **255**

### Numeric Types:

**int**: Represents signed integers, it can store positive and negative numbers. The size of int depends on the architecture of the underlying machine.

**uint**: Represents unsigned integers, it can store only positive numbers. Similar to int, the size depends on the architecture.

**int8, int16, int32, int64**: Signed integers of specific bit sizes.

**byte/uint8, uint16, uint32, uint64**: Unsigned integers of specific bit sizes.

**float32, float64**: Floating-point numbers with single and double precision, respectively.

### String Type:

**string**: Represents a sequence of characters.

### Boolean Type:

**bool**: Represents boolean values, either true or false.

if you want to the variable set a type for you, you can declare the variable without the type, there's two ways for doing this:

```go
// with var word
var test = "bla bla bla"

// without var word
number := 5
```

if you want to see what type is that variable you can do this:

```go
num:=5
fmt.Printf("%T", num)
// output: int
```
