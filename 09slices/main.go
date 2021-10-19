package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{
		"Apple",
		"Mango",
		"Peach",
	}
	fmt.Printf("Type of fruit list is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana") 
	fmt.Println("Fruit list is:", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
}