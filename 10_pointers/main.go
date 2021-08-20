package main

import "fmt"

func main() {
	// Used to point you to a memory address of the value
	a := 5

	// Used &a to use a pointer
	b := &a

	fmt.Println(a, b)

	// * represents a pointer i.e *int
	fmt.Printf("%T\n", b)

	//  Use * to read val from address,
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
