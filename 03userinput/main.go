package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	//comma ok || error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanx for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}