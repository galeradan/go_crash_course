package main

import "fmt"

func main() {

	// Using Var
	var name = "Dan Erick"
	var age int32 = 18
	const isCool = true

	// shorthand
	lastName := "Galera"
	size := 1.3

	fmt.Println(name, age, isCool, lastName, size)
	fmt.Printf("%T\n", size)
}
