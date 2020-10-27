package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	isStarted := startWebServer(port)
	fmt.Println(isStarted)

	err := startWebServer2(port + 1)
	fmt.Println(err)

	_, err3 := startWebServer3(port + 2)
	fmt.Println(err3)
}

func startWebServer(port int) bool {
	fmt.Println("Starting server...")
	// do something
	fmt.Println("Server started on port", port)
	return true
}

func startWebServer2(port int) error {
	fmt.Println("Starting server2...")
	// do something
	fmt.Println("Server2 started on port", port)
	return errors.New("Something went wrong!!!!!!")
}

func startWebServer3(port int) (int, error) {
	fmt.Println("Starting server3...")
	// do something
	fmt.Println("Server3 started on port", port)
	return port, nil
}
