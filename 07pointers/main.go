package main

import (
	"fmt"
)

func main() {
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 25
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is ", myNumber)
}


