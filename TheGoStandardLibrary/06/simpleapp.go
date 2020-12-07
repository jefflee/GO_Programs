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
	timeTrack(start, "Open customer list")

	// create an output file
	outfile, err := os.Create("outfile.csv")
	check(err)
	defer outfile.Close()
	timeTrack(start, "Create outfile")

	// scan the customer list
	scanner := bufio.NewScanner(custlist)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), ",")
		outfile.WriteString(names[1] + "," + names[2] + "\n" )
		outfile.Sync()
	}
	timeTrack(start, "scanned file")
	check(scanner.Err())

	defer timeTrack(start, "Closed Program")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
