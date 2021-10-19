package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Lichi"
	fruitList[2] = "Peach"
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list length is: ", len(fruitList))

	var vegList = [5]string{
		"potato",
		"beans",
		"mushroom",
	}
	fmt.Println("Veggie list is: ", vegList)
	fmt.Println("Veggie list length is: ", len(vegList))
}