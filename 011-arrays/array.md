Arrays in Go have this syntax:

```go
  var arr [5]int
	fmt.Println(arr)
  // [0 0 0 0 0]
```

## to declare an array we add **[]** before the type of the variable, if we put a number inside the [], it will create an array with a fixed length

to access a element of the array we can do like this:

```go
  var arr [5]int
	fmt.Println(arr[0])
  // 0
```

## it just need to put the index of the element which we wanna access, in this case was the first one, so we put [0], if we wanna change the value we can do `arr[0] = 100`

to create an array with the short variable declaration we do like this:

```go
  arr := [3]int{4,5,6}
	fmt.Println(arr)
  // [4, 5, 6]
```
in this case we create an array of 3 length and add the values 4, 5 and 6 to the positions

---
to get the length of the array we can use the **len()** function
```go
  arr := [3]int{4,5,6}
	fmt.Println(len(arr))
  // 3
```