package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("print days in array: ", days)

	for d := 0; d < len(days); d++ {
		fmt.Println("using for loop: ", days[d])
	}

	for i := range days {
		fmt.Println("using range: ", days[i])
	}

	for index, days := range days {
		fmt.Printf("index is %v and value is %v \n", index, days)
	}

	// break
	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 5 {
			break
		}
		fmt.Println("Value is (break): ", rougueValue)
		rougueValue++
	}

	// continue
	value1 := 1
	for value1 < 10 {
		if value1 == 5 {
			value1++
			continue
		}
		fmt.Println("Value is (continue) : ", value1)
		value1++
	}

	//

	value2 := 2
	for value2 < 5 {
		if value2 == 2 {
			fmt.Println("labelState:")
			// value2++
			goto labelStatement
		}
		value2++
		break
	}

labelStatement:
	fmt.Println("Jumping At:  https://github.com/Dhruv9099/golang-learning")

}
