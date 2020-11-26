package main

import (
	"runtime"
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	text, _ := reader.ReadString('\n')
	fmt.Println("Hello %v", text)
	fmt.Printf("We are using Go %v running in %s", runtime.Version(), runtime.GOOS)
}
