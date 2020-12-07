package main

import (
	"fmt"
	"time"
)

func main() {
	// Reference time:
	// Mon Jan 2 15:04:05 MST 2006
	//t := time.Now()
	//fmt.Println(t.Format(time.ANSIC))
	//fmt.Println(t.Format(time.RFC3339))
	//fmt.Println(t.Format(time.UnixDate))
	//fmt.Println(t.Format(time.Kitchen))

	//fmt.Println(t.Format("Mon Jan 2 15:04:05 MST 2006"))
	//fmt.Println(t.Format("Monday, Jan 2, in the year 2006"))

	startDate := time.Date(
		2018, 07, 04, 9, 00, 00, 0, time.UTC)

	fmt.Println(startDate)
	fmt.Println(startDate.Format("Monday,  Jan 2 2006 at 15:04:05"))

}
