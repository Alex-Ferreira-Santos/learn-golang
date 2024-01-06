in Go we have only the for loop, but the for loop has different syntax

the default syntax of for loops are like this:
```go
for y := 0; y <= 5; y++ {
		fmt.Println()
		fmt.Printf("current y value %d", y)
		x++
	}
```
where the first arg is the declaration of the variable, the second one the condition and the third one what's gonna do if the block ended and the codition is still false, in this case we increase the value of y adding 1

---

the second syntax is like a while in other languages
```go
x := 3
for x < 5 {
	fmt.Println()
	fmt.Printf("current x value %d", x)
	x++
}
```

is just the **for** word with the condition, if the condition is false the for will keep running the block of code until the condition becomes true

we can create a infinite loop if we don't pass any conditions to the for:

```go
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
```

in this case we have the variable **a** with 0 value, it enters the for loop and validate two ifs, the first one check if the **a** value is **0**, in this case it will print a message and **continue** the loop, continue is a key word to skip the rest of the block of code and start a new loop, and the second statement validates if the **a** is bigger than **2**, if it is so will print a message and stop the loop with the **break** word