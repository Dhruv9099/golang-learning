package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang.")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Orange"
	fruitList[2] = "Apple"

	fmt.Println("Fruit list:", fruitList)
	fmt.Println("Fruit list:", len(fruitList))

	var vegetables = [3]string{"Tomato", "LadyFinger", "Potato"}
	fmt.Println("Vegetables list is: ", vegetables)
	fmt.Println("Vegetables list is: ", len(vegetables))

}
