Slices are a part of the an array, the difference between the array and slices is that the array has the same length and capacity, meanwhile the slice has a unique length and a capacity of the remaining of items to reach the end of the original array, example:

```go
	var x [5]int = [5]int{1,2,3,4,5}
	var s []int = x[2:4]
	fmt.Println(s)
  // output: [3,4]
  // capacity: 3
  // length: 2
```
in this case we make a slice from an array of five elements, the slice begins at the position 2 of original array and ends at position 4, the length is 2 becase the slice has values [3,4] but the capacity is 3 because we start the slice at position 2, the original array has a length of 5, so 5 - 2 = 3

---

if we don't have a value inside the slice and just put **:** it will clone the original array
```go
	var x [5]int = [5]int{1,2,3,4,5}
	var s []int = x[:]
	fmt.Println(s)
  // output: [1,2,3,4,5]
  // capacity: 5
  // length: 5
```

---

if we just passes the number before the **:**, it will create a slice from that position to the end of the original array:
```go
	var x [5]int = [5]int{1,2,3,4,5}
	var s []int = x[1:]
	fmt.Println(s)
  // output: [2,3,4,5]
  // capacity: 4
  // length: 4
```
same thing if we put only the number after the **:**:

```go
	var x [5]int = [5]int{1,2,3,4,5}
	var s []int = x[:2]
	fmt.Println(s)
  // output: [1,2]
  // capacity: 5
  // length: 2
```

---

to get the capacity of an array or slice we use the `cap()` function

---

it posible to create a slice without having to create a previous array:

```go
var slice []int = []int{1,2,3,4,5}
```
in this case we created a slice with length and capacity of 5, what's happens here is that the slice is created referring a array with the length of the elements

---

to add a value to slice we use the `append()` function:
```go
var slice []int = []int{1,2,3,4,5}
newSlice := append(slice, 10)
fmt.Print(slice)
// [1,2,3,4,5]
fmt.Print(newSlice)
// [1,2,3,4,5,10]
```

in this case we add the number 10 at the last position of the slice, but the append will not add to the original slice, it actually creates a new slice with the value in the end

---

another way to create a slice is creating through the `make()` function
```go
	var slice = make([]int, 5)
	fmt.Printf("slice value %d, length: %d, capacity: %d", slice, len(slice), cap(slice))
  // output: slice value [0 0 0 0 0], length: 5, capacity: 5
```
