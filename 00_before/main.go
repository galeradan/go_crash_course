package main

import (
	"fmt"
)

func main() {
	// --- short hand variable declaration
	// isActive := true

	// Test("Print from a separate file and function")

	// --- IF-ELSE EXAMPLE
	// fmt.Println("--- FOR EXAMPLE")
	// if isActive == true {
	// 	fmt.Println("Hello, It's Active")
	// 	fmt.Println(Addition(13, 12))

	// 	// for loop example
	// 	for i := 0; i <= 3; i++ {
	// 		fmt.Printf("%v\n", Addition(1, i))
	// 	}

	// } else {
	// 	fmt.Println("Nothing is active")
	// }

	// fmt.Println("\n\n")

	// --- ARRAY EXAMPLE
	// fmt.Println("--- ARRAY EXAMPLE")
	// Array()

	// fmt.Println("\n\n")

	// --- SLICES EXAMPLE
	// fmt.Println("--- SLICES EXAMPLE")
	// Slices()
	// fmt.Println("\n\n")

	// --- IOTA and Bit Shift EXAMPLE
	// fmt.Println("--- IOTA Const EXAMPLE")
	// Assigning a role to a user, e.g a Moderator that can see archives and deleted
	// var userOneRole byte = IsMember | CanDelete

	// ArchivePost(userOneRole)

	// --- FizzBuzz
	fmt.Println("--- FizzBuzz EXAMPLE")
	FizzBuzz(1, 5)

}
