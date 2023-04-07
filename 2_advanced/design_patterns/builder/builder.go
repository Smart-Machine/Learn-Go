package main

import "log"

type human struct {
	age      int
	height   float64
	eyeColor string
}

// Empty constructors

func newHuman() human {
	return human{}
}

func newHumanWithFields(age int, height float64, eyeColor string) human {
	return human{
		age:      age,
		height:   height,
		eyeColor: eyeColor,
	}
}

// Builders

func (h human) withAge(a int) human {
	h.age = a
	return h
}

func (h human) withHeight(y float64) human {
	h.height = y
	return h
}

func (h human) withEyeColor(c string) human {
	h.eyeColor = c
	return h
}

// Reset

func (h human) reset() human {
	h = newHuman()
	return h
}

// Pre-defined builds

func giant() human {
	return newHuman().withHeight(2.5)
}

func main() {
	g := giant().withAge(1000).withEyeColor("blue")
	log.Printf("%+v\n", g)

	h := newHuman().withAge(23).withHeight(1.8).withEyeColor("brown")
	log.Printf("%+v\n", h)
}
