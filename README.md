# Learning Go #

This is a cheatsheet for future me about Go.

## Contents ##
* Basics
	* [Hello World](#hello-world)
	* [Packages and Imports](#packages-and-imports)
	* [Named return value](#named-return-value)
	* [Variables](#variables)
	* [Type conversions](#type-conversions)
	* [Constants](#constants)
	* [For loop](#for-loop)
	* [If statement](#if-statement)
	* [Switch statement](#switch-statement)
	* [Defer statement](#defer-statement)
	* [Pointers](#pointers)
	* [Structs](#structs)
	* [Pointers to structs](#pointers-to-structs)
	* [Struct Literals](#struct-literals)
	* [Arrays](#arrays)
	* [Slices](#slices)
	* [Slice length and capacity](#slice-length-and-capacity)
	* [Creating a slice with make](#creating-a-slice-with-make)
	* [Appending to a slice](#appending-to-a-slice)
	* [Range](#range)
	* [Maps](#maps)
	
### Hello World ###

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```



### Packages and Imports ###

Every go file has to be assigned to a package. This helps group files, and share their content between themselves. For example, if you want to use a function from a different file in your `main` file, you have to put them in the same package. 

*main.go*
```go
package main

import "fmt"

func main() {
	fmt.Println(hello())
}
```

*function.go*
```go
package main

func hello() string {
	return "Hello"
}
```

```bash
$ go run *.go
Hello
```

Imports are usually written as follows
```go
import "fmt"
```

```go
import "fmt"
import "math"
```

```go
import (
	"fmt"
	"math"
)
```



### Named return value ###

Go's return values may be named. If so, they are treated as variables defined at the top of the function.
```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

```bash
7 10
```



### Variables ###

It is considered a good practice to use the `var` keyword for declaring variables outside the scope of a function, or at the beginning of a big function. For example:
```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```
Also, you can initialize the variable as you declare it. 
```go
var i, j int = 1, 2
```
Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.
```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

```bash
1 2 3 true false no!
```



### Type conversions ###
The expression `T(v)` converts the value `v` to the type `T`.
Some numeric conversions:
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
Or, put more simply:
```go
i := 42
f := float64(i)
u := uint(f)
```
Unlike in C, in Go assignment between items of different type requires an explicit conversion. 



### Constants ###
Constants are declared like variables, but with the `const` keyword. Constants can be character, string, boolean, or numeric values. Constants cannot be declared using the `:=` syntax.
```go
const Truth = true
fmt.Println("Go rules?", Truth)
```

```bash
Go rules? true
```



### For loop ###
Go has only one looping construct, the `for` loop. The basic `for` loop has three components separated by semicolons:

* **the init statement**: executed before the first iteration
* **the condition expression**: evaluated before every iteration
* **the post statement**: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the `for` statement.

**Note**: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the `for` statement and the braces `{ }` are always required.

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

```bash
45
```

C's `while` is spelled `for` in Go.

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

```bash
1024
```

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

```go
package main

func main() {
	for {
	}
}
```



### If statement ###

Go's `if` statements are like its `for` loops; the expression need not be surrounded by parentheses `( )` but the braces `{ }` are required.

```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

```bash
1.4142135623730951 2i
```

Like `for`, the `if` statement can start with a short statement to execute before the condition. Variables declared by the statement are only in scope until the end of the `if`.

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim   //if you will try to return here v, it will fail.
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

```bash
9 20
```

`If` and `else` example:

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

```bash
27 >= 20
9 20
```



### Switch statement ###

A `switch` statement is a shorter way to write a sequence of `if - else` statements. It runs the first case whose value is equal to the condition expression. 
Go's `switch` is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's `switch` cases need not be constants, and the values involved need not be integers.

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

```bash
Go runs on Linux.
```

The `switch` construction can have no condition.

```go 
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```



### Defer statement ###

A defer statement defers the execution of a function until the surrounding function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

```bash
hello
world
```

You can stack defer. Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order. More about this [here](https://go.dev/blog/defer-panic-and-recover).

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

```bash
counting
done
9
8
7
6
5
4
3
2
1
0
```



### Pointers ###

Go has pointers. A pointer holds the memory address of a value. The type `*T` is a pointer to a `T` value. Its zero value is `nil`.
```go 
var p *int
```
The `&` operator generates a pointer to its operand.
```go
i := 42
p = &i
```
The `*` operator denotes the pointer's underlying value.
```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```
**Note**: Unlike C, Go has no pointer arithmetic.

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

```bash
42
21
73
```



### Structs ###

A `struct` is a collection of fields. 

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

Struct fields are accessed using a dot.

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```



### Pointers to structs ###

Struct fields can be accessed through a struct pointer.

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

To access the field `X` of a struct when we have the struct pointer `p` we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.



### Struct Literals ###

A struct literal denotes a newly allocated struct value by listing the values of its fields. You can list just a subset of fields by using the `Name:` syntax. 

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```

```bash
{1 2} &{1 2} {1 0} {0 0}
```



### Arrays ###

The type `[n]T` is an array of `n` values of type `T`.

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

```bash
Hello World
[Hello World]
[2 3 5 7 11 13]
```



### Slices ###

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays. The type `[]T` is a slice with elements of type `T`.
A slice is formed by specifying two indices, a low and high bound, separated by a colon:
```cmd
a[low : high]
```
This selects a half-open range which includes the first element, but excludes the last one.

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

```bash
[3 5 7]
```

A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

```bash
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
```

```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

```bash
[2 3 5 7 11 13]
[true false true true false true]
[{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
```

When slicing, you may omit the high or low bounds to use their defaults instead. The default is zero for the low bound and the length of the slice for the high bound. 

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
```

```bash
[3 5 7]
[3 5]
[5]
```

Slices can contain any type, including other slices.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

```bash
X _ X
O _ X
_ _ O
```


### Slice length and capacity ###

A slice has both a *length* and a *capacity*. The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`. 
You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

```bash
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]
```



### Creating a slice with make ###

Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.
The `make` function allocates a zeroed array and returns a slice that refers to that array:

```go 
a := make([]int, 5)  // len(a)=5
```

To specify a capacity, pass a third argument to `make`:

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```

```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

```bash
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]
```



### Appending to a slice ###

It is common to append new elements to a slice, and so Go provides a built-in `append` function.

```go 
func append(s []T, vs ...T) []T
```

The first parameter `s` of append is a slice of type `T`, and the rest are `T` values to append to the slice.
The resulting value of `append` is a slice containing all the elements of the original slice plus the provided values. If the backing array of `s` is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

```bash
len=0 cap=0 []
len=1 cap=1 [0]
len=2 cap=2 [0 1]
len=5 cap=6 [0 1 2 3 4]
```



### Range ###

The `range` form of the `for` loop iterates over a slice or map.
When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

```bash
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
```

You can skip the index or value by assigning to `_`.

```go
for i, _ := range pow
for _, value := range pow
```

If you only want the index, you can omit the second variable.

```go
for i := range pow
```

```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

```bash
1
2
4
8
16
32
64
128
256
512
```



### Maps ###

A map maps keys to values. The `make` function returns a map of the given type, initialized and ready for use.

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

```bash
{40.68433 -74.39967}
```

Map literals are like struct literals, but the keys are required.

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
```

```bash
map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
```

If the top-level type is just a type name, you can omit it from the elements of the literal.

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```

```bash
map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
```




