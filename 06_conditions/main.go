package main

import "fmt"

func main() {
	x := 5
	y := 10

	//  If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "dark"

	if color == "dark" {
		fmt.Println("color is dark")
	} else if color == "light" {
		fmt.Println("color is light")
	} else {
		fmt.Println("color is NOT light or dark")
	}

	// Switch
	switch color {
	case "dark":
		fmt.Println("color is dark")
	case "light":
		fmt.Println("color is light")
	default:
		fmt.Println("color is NOT light or dark")
	}

}
