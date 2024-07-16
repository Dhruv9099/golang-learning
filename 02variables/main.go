package main

import "fmt"

//  check the documentation go programming language specification > Types> Numeric Types
func main() {
	fmt.Println("It is 02Variables.go file. ")
	var username string = "Dhruv Maisuria"
	// fp
	fmt.Println(username)

	// When you use %T, it will display the type of the variable passed to fmt.Printf.
	fmt.Printf("Variable is of type: %T \n", username)

	// bool
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	// ff
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	fmt.Printf("%T \n", isLoggedIn)

	// integer
	var smallVal int = 9099
	fmt.Println(smallVal)
	// ff
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// Float
	var smallFloatVal float32 = 9.9099826549
	fmt.Println(smallFloatVal)
	// ff
	fmt.Printf("Variable is of type: %T \n", smallFloatVal)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of types: %T \n", anotherVariable)

	// implicit type
	var website = "www.example.com"
	fmt.Println(website)
	// website =3  will not accept because lexter save as string first.

	// no var style  allowed only in function, not outside of function

	numberOfUser := 50
	fmt.Println("number of users: ", numberOfUser)
	fmt.Printf("Type of number of user: %T \n", numberOfUser)
	fmt.Println(LoginToken)
	fmt.Printf("type of LoginToken is : %T \n", LoginToken)
}

const LoginToken string = "JayShreeRam " //public
