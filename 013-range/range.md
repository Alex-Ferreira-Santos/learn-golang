In Go we can use the **`range`** to iterate an array or slice without using the default for loop

```go
var arr []int = []int{1,2,34,2,123,4532,65}
for index, element := range arr {
	fmt.Printf("index: %d, element: %d\n", index, element)
}
/*
output:
index: 0, element: 1
index: 1, element: 2
index: 2, element: 34
index: 3, element: 2
index: 4, element: 123
index: 5, element: 4532
index: 6, element: 65
*/
```
in this case we created a slice, then we created a for loop but instead of using the default condition we use **`range`** to iterate the array, the range arr returns two values, the index of the array and the element that is stored in that index position, in case we don't want the index we can replace it with a underscore _