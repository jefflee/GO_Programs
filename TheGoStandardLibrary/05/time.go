package main

import (
	"fmt"
	"time"
)

func main() {

	// Reference time:
	// Mon Jan 2 15:04:05 MST 2006

	/*
	first := time.Now()
	fmt.Printf("It is currently %v\n", first.Format("15:04:05"))
	time.Sleep(1000000000)
	time.Sleep(1 * time.Second)
	second := time.Now()
	fmt.Printf("It is now %v\n", second.Format("15:04:05"))
	*/

	/*
	today := time.Now()
	fmt.Printf("It is currently %v\n", today.Format("Monday, Jan 2 2006"))
	startDate := time.Date(2018, 07, 04, 9, 00, 00, 0, time.UTC)
	elapsed := time.Since(startDate)

	fmt.Printf("Hours: %.0f Minutes: %.0f Seconds: %.0f\n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())
	*/

/*
	today := time.Now()
	future := today.Add(6 * time.Hour)

	fmt.Printf("six months ago it was %v\n", future.Format("Monday, January 2"))
*/

	bedtime := time.Date(2020, 6, 6, 23, 0, 0, 0, time.Local)

	fmt.Printf("There is %.0f hours until bed time\n", time.Until(bedtime).Hours())


}

