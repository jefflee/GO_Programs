package mylib

import "fmt"

func messageWriter(greeting, name string) string {
	message := fmt.Sprintf("%v, %v", greeting, name)
	return message
}
