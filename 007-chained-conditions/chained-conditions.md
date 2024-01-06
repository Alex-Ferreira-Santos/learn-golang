Chained conditions:

! = Not
|| = Or
&& = And

```go
num1:= 1
num2:= 2

fmt.Print(!(num1 > num2))
// output: true

fmt.Print(num1 > 1 && num2 > 1)
// output: false

fmt.Print(num1 > 1 || num2 > 1)
// output: true

```
