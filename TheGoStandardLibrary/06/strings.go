package main

import "fmt"

func main() {
	ourString := "It is awesome!!"

	// print each byte of a string.
	for i := 0; i < len(ourString); i++ {
		fmt.Printf("%x ", ourString[i])
	}
	fmt.Println()

	anotherStringVariable := "\x49\x74\x20\x69\x73\x20\x61\x77\x65\x73\x6f\x6d\x65\x21\x21"
	fmt.Printf("%v \n", anotherStringVariable)

	// print each character
	for i := 0; i < len(ourString); i++ {
		fmt.Printf("%q \n", ourString[i])
	}
	fmt.Println()

	newString := "This is a string!"

	// print 115 -> It's byte value
	fmt.Println(newString[3])

	// print this
	fmt.Println(newString[0:5])
}
