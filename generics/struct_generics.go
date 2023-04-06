package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}

func main() {
	u := User[string]{
		ID:   0,
		Name: "Calin",
		Data: "some data for the user",
	}
	fmt.Printf("user: %+v\n", u)
}
