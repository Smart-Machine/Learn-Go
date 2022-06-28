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
The expression T(v) converts the value v to the type T.
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









