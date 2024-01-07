Maps in Go are created with the `map` keyword

```go
var mp map[string]int = map[string]int{
		"apple": 1,
		"pear": 6,
		"orange": 10,
	}
fmt.Println(mp)
// output: map[apple:1 orange:10 pear:6]
```

Maps are a type of structure that store values with a **key** to reference a **value**
Maps don't store the values in order, so if the order it's important, don't use it!

---

There are more ways to create a `map` in Go:

```go
map2 := make(map[string]int)
fmt.Println(map2)
// output:  map[]
```

---

to access a value inside a map we use the key of the element

```go
var mp map[string]int = map[string]int{
		"apple": 1
	}
fmt.Print(mp["apple"])
// output: 1
```

you can change the value using the key too

```go
var mp map[string]int = map[string]int{
		"apple": 1
	}
  mp["apple"] = 100
fmt.Print(mp["apple"])
// output: 100
```
to add a new value into the map you just need to something like: `mp["new key"] = "new value"`

to delete a value from the map you need to use the `delete()` function:
```go
var mp map[string]int = map[string]int{
		"apple": 1
}
delete(mp, "apple")
fmt.Print(mp["apple"])
// output: 0
```
