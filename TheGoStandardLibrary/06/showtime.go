package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()


	Year := t.Year()
	Month := t.Month()
	Day := t.Day()

	fmt.Printf("Today is %d/%d/%v", Month, Day, Year)

}
