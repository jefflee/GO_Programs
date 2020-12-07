package main

import "time"

func main() {
	then := time.Date(2020, 4, 1, 05, 00, 00, 0, time.UTC)
	now := time.Now()
	diff := now.Sub(then)
	print(diff.Hours())
}
