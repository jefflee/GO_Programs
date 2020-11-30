package main

import (
	"fmt"
)

func main() {
	var firstNumber int
	var secondNumber int
	fmt.Println("What two numbers would you like to add?")
	fmt.Scanf("%d %d", &firstNumber, &secondNumber)
	fmt.Printf("Your total is %d\n", firstNumber + secondNumber)
}
