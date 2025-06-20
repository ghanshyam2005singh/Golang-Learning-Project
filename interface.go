package main

import "fmt"

// Define the interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle implements Shape
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %v, Perimeter: %v\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{width: 4, height: 3}
	circle := Circle{radius: 2}

	printShapeInfo(rect)
	printShapeInfo(circle)
}
