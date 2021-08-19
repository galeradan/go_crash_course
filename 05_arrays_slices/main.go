package main

import "fmt"

func main() {

	// Declare Then Assign
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(fruitArr[0])

	// Declare And Assign
	animeArr := [2]string{"Tensei", "Tempest"}

	fmt.Println(animeArr)
	fmt.Println(animeArr[0])
	fmt.Println(animeArr[1])

	fruitSlice := [3]string{"Apple", "Orange"}

	fmt.Println(fruitSlice)

	// Length
	fmt.Println(len(fruitSlice))

	// Range
	fmt.Println(cap(fruitSlice))

}
