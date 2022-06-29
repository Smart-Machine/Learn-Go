# Learning Go #

This is a cheatsheet for future me about Go.



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

```cmd
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

```cmd
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

```cmd
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

```cmd
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

```cmd
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

```cmd
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

```cmd
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

```cmd
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

```cmd
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

```cmd
Go runs on Linux.
```





