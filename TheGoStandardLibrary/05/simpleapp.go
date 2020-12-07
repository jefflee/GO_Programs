package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	args := os.Args
	// open the customer list
	custlist, err := os.Open(string(args[1]))
	check(err)
	defer custlist.Close()

	writeTime(start, "Opened Customer List")

	// create an output file
	outfile, err := os.Create("outfile.csv")
	check(err)
	defer outfile.Close()

	writeTime(start, "Created outfile")

	// scan the customer list
	scanner := bufio.NewScanner(custlist)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), ",")
		outfile.WriteString(names[1] + "," + names[2] + "\n")
	}
	check(scanner.Err())
	writeTime(start, "Wrote data to outfile")
	defer writeTime(start, "Application Finished")

}

func writeTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
