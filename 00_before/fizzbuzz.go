package main

import "fmt"

// FizzBuzz is a classic programming challenge
func FizzBuzz(min int, max int) {
	for i := min; i <= max; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}

	// count := 0
	// for a := min; a <= max; a++ {
	// 	for b := a; b <= max; b++ {
	// 		n := a * b

	// 		s := fmt.Sprintf("%d", n)
	// 		if s[0] == s[len(s)-1] {
	// 			count++
	// 		}
	// 	}
	// }

	// fmt.Println(count)
}
