package main

import "fmt"

// Slices func showcases how to use slices in go
func Slices() {

	// declaration
	slice := []int{1, 2, 3}
	fmt.Println("Slices", slice)

	// built in function
	fmt.Printf("Lenght is %v\n", len(slice))
	fmt.Printf("Capacity is %v\n", cap(slice))

	// slices are naturally reference, no need to use &
	sliceRef := slice
	slice[1] = 5

	// both arrays will change
	fmt.Println("slice: ", slice)
	fmt.Println("sliceRef: ", sliceRef)

}
