package main


import (
	"fmt"
)

func main() {

	var firstNumber int
	println("What's the first number?")
	fmt.Scanln("%d", &firstNumber)

	var secondNumber int
	println("What's the second number?")
	fmt.Scanln("%d", &secondNumber)

	fmt.Printf("Answer is: %d", firstNumber + secondNumber)
}
