package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{2}
	rectangle := Rectangle{2, 5}

	fmt.Println("Hello World", getArea(circle), getArea(rectangle))
}
