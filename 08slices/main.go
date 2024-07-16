package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices  section in go")

	// var fruitList = []string{}
	var fruitList = []string{"apple", "banana"}
	fmt.Println(fruitList)
	fmt.Printf("Type of fruitList is : %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Cherry", "berry")
	fmt.Println("FruitList: ", fruitList)

	fruitList = (fruitList[1:])
	fmt.Println("FruitList[1:]", fruitList)

	fruitList = (fruitList[:3])
	fmt.Println("FruitList[:3]", fruitList)

	fruitList = (fruitList[1:1])
	fmt.Println("FruitList[1:1]", fruitList)

	fruitList = (fruitList[1:2])
	fmt.Println("FruitList[1:2]", fruitList)
	fmt.Printf("\n")
	highScores := make([]int, 4)
	highScores[0] = 523
	highScores[1] = 856
	highScores[2] = 956
	highScores[3] = 999

	fmt.Println("highScores: ", highScores)
	//append in array
	highScores = append(highScores, 555, 362, 989)
	fmt.Println("append highScores: ", highScores)

	// sort
	fmt.Println("sorted array:", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("sorted array:", highScores)
	fmt.Println("sorted array:", sort.IntsAreSorted(highScores))
	fmt.Printf("\n")

	// how to remove a value from slices based on index
	var courses = []string{"reactJs", "javascript", "python", "go", "ruby", "python"}
	fmt.Println("course: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("", courses)
}
