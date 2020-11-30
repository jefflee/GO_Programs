package main

import (
	"fmt"
)

func main() {
	fmt.Print("What is your name? ")
	var firstname string
	var lastname string
	fmt.Scanln(&firstname, &lastname)
	fmt.Printf("Hello %s %s nice to meet you!\n", firstname, lastname)
}
