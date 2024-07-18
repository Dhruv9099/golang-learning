package main

import "fmt"

// func main() {}

func main() {
	fmt.Println("Welcome to function in golang")
	greeting2()
	greeting1()
	result := add(3, 5)
	fmt.Println("Result using add() function: ", result)

	proResult, MyMessage := proAdd(3, 5, 8, 8, 9)
	// proResult, _ := proAdd(3, 5, 8, 8, 9)
	fmt.Println("Result using proAdd() function: ", proResult)
	fmt.Println("Result using proAdd() myMessage ", MyMessage)
}

func greeting1() {
	fmt.Println("Hello from greeting1() function return in main() function.")
}

func greeting2() {
	fmt.Println("Hello from greeting2() function return in main() function.")
}

// function signature
func add(val1 int, val2 int) int {
	return val1 + val2
}

//
func proAdd(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "this is proAdd function you can add unlimited sum"
}
