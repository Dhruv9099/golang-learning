package main

import "fmt"

func main() {
	fmt.Println("welcome to If else in golang.")

	// example 1
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User."
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		// fmt.Printf("Exactly %v login count.", loginCount)
		result = "Exactly 10 login count."
	}
	fmt.Println(result)

	// example 2
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
	// example 3
	if num := 10; num < 10 {
		fmt.Println("Number is less than 10")
	} else if num == 10 {
		fmt.Println("Number is equals to 10")
	} else {
		fmt.Println("Number is more than 10")

	}
}
