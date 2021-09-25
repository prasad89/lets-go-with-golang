package main

import "fmt"

const LoginToken string = "This is string" //Public

func main() {
	var username string = "Vivek"
	fmt.Println(username)
	fmt.Printf("varible is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("varible is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("varible is of type: %T \n", smallVal)

	var smallFloat float64 = 255.255255255255255
	fmt.Println(smallFloat)
	fmt.Printf("varible is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("varible is of type: %T \n", anotherVariable)

	//implicit type
	var website = "www.google.com"
	fmt.Println(website)

	//no var style
	numberOfUser := 3000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("varible is of type: %T \n", LoginToken)
}