package main

import (
	"fmt"
	"reflect"

)

// Main function
func main() {

	v := reflect.TypeOf("123")

	// use of SliceOfmethod
	fmt.Println(reflect.SliceOf(v))

}