package main

import "fmt"

func main() {
	fmt.Println("Welcome to  pointer  topic in GO.")

	// var one int = 2  in this if you  don't  give value

	// Pointer Declaration:
	var ptr *int

	// Printing the Pointer Value:
	fmt.Println("Value of pointer is: ", ptr)
	// output:- Value of pointer is:  <nil>

	// Assigning a Value to a Variable:
	myNumber := 24

	// Pointer Assignment:
	var pointer = &myNumber

	// Printing the Pointer Address:
	fmt.Println("Values of  actual Pointer is : ", pointer)
	// Values of  actual Pointer is :  0xc00000a0b8

	// Dereferencing the Pointer:
	fmt.Println("Values of  actual Pointer is : ", *pointer)
	// Values of  actual Pointer is :  24

	// Modifying the Value through the Pointer:
	*pointer = *pointer * 2
	fmt.Println("New Values is :", myNumber)
}

// Pointer: A variable that stores the memory address of another variable.
// 	Pointers are used for referencing and dereferencing memory locations.

// & Operator: Used to get the address of a variable.

// *** Operator**: Used to access the value stored at the address a pointer is pointing to.
