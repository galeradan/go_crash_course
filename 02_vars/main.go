package main

import "fmt"

func main() {

	// Using Var
	var name = "Dan Erick"
	var age int32 = 18
	const isCool = true

	// shorthand
	last_name := "Galera"

	fmt.Println(name, age, isCool, last_name)
	fmt.Printf("%T\n", isCool)
}
