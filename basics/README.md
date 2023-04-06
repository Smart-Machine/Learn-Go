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
	* [Mutating Maps](#mutating-maps)
	* [Function values](#function-values)
	* [Function closures](#function-closures)
* Methods and interfaces
	* [Methods](#methods)
	* [Methods continued](#methods-continued)
	* [Pointer receivers](#pointer-receivers)
	* [Methods and pointer indirection](#methods-and-pointer-indirection)
	* [Choosing a value or pointer receiver](#choosing-a-value-or-pointer-receiver)
	* [Interfaces are implemented implicitly](#interfaces-are-implemented-implicitly)
	* [Interface values](#interface-values)
	* [The empty interface](#the-empty-interface)
	* [Type assertions](#type-assertions)
	* [Type switches](#type-switches)
	* [Stringers](#stringers)
	* [Errors](#errors)
* Generics
	* [Type parameters](#type-parameters)
	* [Generic types](#generic-types)
* Concurrency 
	* [Goroutines](#goroutines)
	* [Channels](#channels)
	* [Buffered Channels](#buffered-channels)
	* [Range and Close](#range-and-close)
	* [Select](#select)
	* [Default Selection](#default-selection)
	
	
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



### Mutating Maps ###

Insert or update an element in map `m`:

```go
m[key] = elem
```

Retrieve an element:

```go
elem = m[key]
```

Delete an element:

```go
delete(m, key)
```

A key is present with a two-value assignment:

```go
elem, ok = m[key]
```

If `key` is in `m`, `ok` is `true`. If not, `ok` is `false`.
If `key` is not in the map, then `elem` is the zero value for the map's element type.

**Note**: If `elem` or `ok` have not yet been declared you could use a short declaration form:

```go
elem, ok := m[key]
```

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
```

```bash
The value: 42
The value: 48
The value: 0
The value: 0 Present? false
```



### Function values ###

Functions are values too. They can be passed around just like other values. Function values may be used as function arguments and return values.

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

```bash
13
5
81
```



### Function closures ###

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the `adder` function returns a closure. Each closure is bound to its own `sum` variable.

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

```bash 
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
```

Let's have some fun with functions. 
Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		s := a + b
		a = b
		b = s
		return s
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

```bash
1
2
3
5
8
13
21
34
55
89
```

### Methods ###

Go does not have classes. However, you can define methods on types.
A method is a function with a special __receiver__ argument.
The receiver appears in its own argument list between the `func` keyword and the method name.
In this example, the Abs method has a receiver of type `Vertex` named `v`.

_methods.go_
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```
_output:_
```
5
```

Remember: a method is just a function with a receiver argument.
Here's Abs written as a regular function with no change in functionality.


### Methods continued ###

You can declare a method on non-struct types, too.
In this example we see a numeric type `MyFloat` with an `Abs` method.
You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).

_methods-continued.go_
```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```

_output:_
```
1.4142135623730951
```



### Pointer receivers ###
You can declare methods with pointer receivers.
This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`.)
For example, the `Scale` method here is defined on `*Vertex`.
Methods with pointer receivers can modify the value to which the receiver points (as `Scale` does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
Try removing the `*` from the declaration of the `Scale` function on line 16 and observe how the program's behavior changes.
With a value receiver, the `Scale` method operates on a copy of the original `Vertex` value. (This is the same behavior as for any other function argument.) The `Scale` method must have a pointer receiver to change the `Vertex` value declared in the `main` function.

_methods-pointer.go_
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

_output:_
```
50
```

### Methods and pointer indirection ###
Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:
```go
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```
while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
```go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```
For the statement `v.Scale(5)`, even though `v` is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale` method has a pointer receiver.

_indirection.go_
```go
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
```

_output:_
```
{60 80} &{96 72}
```


### Choosing a value or pointer receiver ###
There are two reasons to use a pointer receiver.
The first is so that the method can modify the value that its receiver points to.
The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
In this example, both `Scale` and `Abs` are methods with receiver type `*Vertex`, even though the `Abs` method needn't modify its receiver.
In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

_methods-with-pointer-receivers.go_
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

_output:_
```
Before scaling: &{X:3 Y:4}, Abs: 5
After scaling: &{X:15 Y:20}, Abs: 25
```



### Interfaces are implemented implicitly ###
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

_interfaces-are-satisfied-implicitly.go_
```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

_output:_
```
hello
```



### Interface values ###
Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
```
(value, type)
```
An interface value holds a value of a specific underlying concrete type.
Calling a method on an interface value executes the method of the same name on its underlying type.

_interface-value.go_
```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

_output:_
```
(&{Hello}, *main.T)
Hello
(3.141592653589793, main.F)
3.141592653589793
```


### The empty interface ###
The interface type that specifies zero methods is known as the empty interface:
```go
interface{}
```
An empty interface may hold values of any type. (Every type implements at least zero methods.)
Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.

_empty-interface.go_
```go
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

_output:_
```
(<nil>, <nil>)
(42, int)
(hello, string)
```



### Type assertions ###
A _type assertion_ provides access to an interface value's underlying concrete value.
```go
t := i.(T)
```
This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.
If `i` does not hold a `T`, the statement will trigger a panic.
To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
```go
t, ok := i.(T)
```
If `i` holds a `T`, then `t` will be the underlying value and ok will be true. If not, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.

_type-assertions.go_
```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```

_output:_
```
hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64
```



### Type switches ###
A _type switch_ is a construct that permits several type assertions in series.
A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.
```go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```
The declaration in a type switch has the same syntax as a type assertion `i.(T)`, but the specific type `T` is replaced with the keyword `type`.

This switch statement tests whether the interface value `i` holds a value of type `T` or `S`. In each of the `T` and `S` cases, the variable `v` will be of type `T` or `S` respectively and hold the value held by `i`. In the default case (where there is no match), the variable `v` is of the same interface type and value as `i`.

_type-switches.go_
```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

_output:_
```
Twice 21 is 42
"hello" is 5 bytes long
I don't know about type bool!
```



### Stringers ###
One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package.
```go
type Stringer interface {
    String() string
}
```
A `Stringer` is a type that can describe itself as a string. The `fmt` package (and many others) look for this interface to print values.

_stringer.go_
```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```

_output:_
```
Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
```


### Errors ### 
Go programs express error state with `error` values.
The `error` type is a built-in interface similar to `fmt.Stringer`:
```go
type error interface {
    Error() string
}
```
(As with `fmt.Stringer`, the fmt package looks for the `error` interface when printing values.)
Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.
```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```
A nil `error` denotes success; a non-nil `error` denotes failure.

_errors.go_
```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```

_output:_
```
at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work
```



### Type parameters ### 
Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets, before the function's arguments.
```go
func Index[T comparable](s []T, x T) int
```
This declaration means that `s` is a slice of any type `T` that fulfills the built-in constraint `comparable`. `x` is also a value of the same type.

`comparable` is a useful constraint that makes it possible to use the `==` and `!=` operators on values of the type. In this example, we use it to compare a value to all slice elements until a match is found. This `Index` function works for any type that supports comparison.

_index.go_
```go
package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
```

_output:_
```
2
-1
```



### Generic types ###
In addition to generic functions, Go also supports generic types. A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.

This example demonstrates a simple type declaration for a singly-linked list holding any type of value.

As an exercise, add some functionality to this list implementation.

_list.go_
```go
package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}
```



### Goroutines ###
A _goroutine_ is a lightweight thread managed by the Go runtime.
```go
go f(x, y, z)
```
starts a new goroutine running
```
f(x, y, z)
```
The evaluation of `f`, `x`, `y`, and `z` happens in the current goroutine and the execution of `f` happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives. 

_goroutines.go_
```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
```

_output:_
```
world 
hello
hello 
world
world
hello
...
```


### Channels ###
Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.
```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```
Like maps and slices, channels must be created before use:
```go
ch := make(chan int)
```
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables. The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

_channels.go_
```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

_output:_
```
-5 17 12
```


### Buffered Channels ###
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
```go
ch := make(chan int, 100)
```
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.



### Range and Close ###
A sender can `close` a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression after:
```go
v, ok := <-ch
```
`ok` is `false` if there are no more values to receive and the channel is closed. 
The loop for `i := range c` receives values from the channel repeatedly until it is closed.

**Note:** Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

**Another note:** Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a `range` loop.

_range-and-close.go_
```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```

_output:_
```
0
1
1
2
3
5
8
13
21
34
```



### Select ###
The `select` statement lets a goroutine wait on multiple communication operations.

A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

_select.go_
```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

_output:_
```
0
1
1
2
3
5
8
13
21
34
quit
```



### Default Selection ###

The `default` case in a `select` is run if no other case is ready.
Use a `default` case to try a send or receive without blocking:
```go
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

_default-selection.go_
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

_output:_
```
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
BOOM!
```

