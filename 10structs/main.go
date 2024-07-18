package main

import "fmt"

func main() {
	fmt.Println("Welcome to Struct method in go lang. ")

	// no inheritance , no super or parent

	Dhruv := User{
		"DhruvKumar", "dhruv9099@gmail.com", true, 21}
	fmt.Println("Dhruv1:", Dhruv)

	// %v default format
	fmt.Printf("Dhruv2: %v\n", Dhruv)

	//  %+v: Includes field names in structs.
	fmt.Printf("Dhruv3: %+v\n", Dhruv)

	fmt.Printf("person name is %v and Email is %v and Age is %v.", Dhruv.Name, Dhruv.Email, Dhruv.Age)
}

// struct are define as collection of fields each with name and data type
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
