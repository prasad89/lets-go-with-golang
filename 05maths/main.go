package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	//var numberOne int = 2
	//var numberTwo float64 = 4.5
	//fmt.Println("The sum is: ", numberOne + int(numberTwo))

	//random number
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5) + 1)

	//random form crypto
	randomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNum)
}