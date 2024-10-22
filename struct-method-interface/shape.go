package main

import (
	"math"
)

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func Parimeter(rectangle Rectangle) (result float64) {
	return 2 * (rectangle.height + rectangle.width)
}

func (r Rectangle) Area() (result float64) {
	return r.height * r.width
}

func (c Circle) Area() (result float64) {
	return math.Pi * c.radius * c.radius
}
