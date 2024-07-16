package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map in go lang.")

	languages := make(map[string]string)
	languages["Js"] = "JavaScript"
	languages["rb"] = "Ruby"
	languages["go"] = "GoLang"
	languages["py"] = "Python"

	fmt.Println("List of Languages: ", languages)
	fmt.Println("Js shorts for : ", languages["Js"])

	// delete
	delete(languages, "rb")
	fmt.Println("After delete(language,`rb`: )", languages)

	// loop are interesting in golang
	// method 1
	for key, value := range languages {
		fmt.Printf("For key: %v, value is: %v \n", key, value)
	}
	// method 2
	for _, value := range languages {
		fmt.Printf("For KEY: , value is: %v \n", value)
	}
}
