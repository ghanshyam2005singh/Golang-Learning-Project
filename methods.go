package main

import "fmt"

// Define a struct
type Rectangle struct {
	width, height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method with pointer receiver (can modify the struct)
func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 5, height: 3}
	fmt.Println("Area:", rect.Area())

	rect.Scale(2)
	fmt.Println("Scaled Area:", rect.Area())
}
