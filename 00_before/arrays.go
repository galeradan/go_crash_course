package main

import "fmt"

// Array list examples of how to use arrays in go
func Array() {
	// var testArr [3]int one-dimen
	var testArr [3][2][4][1]int // Multi-dimen
	fmt.Println("array: ", testArr)

	// shorthand declaration
	shortArray := [3]int{1, 2, 3}
	fmt.Println("shorthand array: ", shortArray)

	// ... literals if no need to specificy
	literalArray := [...]int{1, 2, 3, 4}
	fmt.Println("literal array: ", literalArray)

	// using an array
	fmt.Println("The value of shortArray[0] is", shortArray[0])

	//knowing the length of the array
	fmt.Println("Lenght of literalArray is", len(literalArray))

	//copying arrays
	arrA := [...]int{3, 2, 1}
	arrB := arrA
	arrB[1] = 4
	// arrA will not change
	fmt.Println("arrA: ", arrA)
	// arrB will be change
	fmt.Println("arrB: ", arrB)

	arrC := [...]int{7, 8, 9}
	arrD := &arrC
	arrD[2] = 5
	// arrC will change as well because arrD is using & which points directly to arrC
	fmt.Println("arrC: ", arrC)
	fmt.Println("arrD: ", arrD)

	// End of Array Example
}
