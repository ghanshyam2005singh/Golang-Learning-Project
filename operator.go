package main

import "fmt"

func main() {
	a := 10
	b := 3

	// Arithmetic
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// Assignment
	a += 5 // a = a + 5
	fmt.Println("a after a += 5:", a)

	// Comparison
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)

	// Logical
	isAdult := true
	hasTicket := false
	fmt.Println("isAdult && hasTicket:", isAdult && hasTicket)
	fmt.Println("isAdult || hasTicket:", isAdult || hasTicket)
	fmt.Println("!isAdult:", !isAdult)

	// Increment/Decrement
	a++
	fmt.Println("a after increment:", a)
	b--
	fmt.Println("b after decrement:", b)
}
