In Go there's mutable and immutable data

```go
var num1 int = 5
num2:= num1
num2 = 7
fmt.Printf("num1: %d, num2: %d \n", num1, num2)
// output: num1: 5, num2: 7

var slice []int = []int{3,4,5}
newSlice := slice
newSlice[0] = 100
fmt.Printf("slice: %d, newSlice: %d",slice, newSlice)
// output: slice: [100 4 5], newSlice: [100 4 5]
```

in this case integer variable are immutable, if you create a new variable **num1** recieving a integer variable **num2** as initial value and then change the new variable **num2**, the **num1** will continue to have the same value as before

but the thing changes when we are talking about slices, if you create **slice1** and then **slice2** recieving the **slice1** and then change a value from **slice2**, it will affect the **slice1**, because slices are mutable and when you creating a new **slice** from another **slice** you are just making a reference from the original slice, same happens if you do this with a map

```go
var mp map[string] int = map[string]int{"hello": 3}
newMap := mp
newMap["hello"] = 300
fmt.Println(newMap, mp)
// output: map[hello:300] map[hello:300]
```

but if we use array instead of slices it won't change both arrays

```go
var array [2]int = [2]int{3,4}
newArr := array
newArr[0] = 500
fmt.Printf("array: %d, newArr: %d\n",array, newArr)
// output: array: [3 4], newArr: [500 4]
```
