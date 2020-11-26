package main

import (
	"flag"
	"fmt"
)

func main() {

	bitsPtr := flag.String("arch", "x86", "CPU Architecture")

	flag.Parse()

	switch *bitsPtr {

	case "x86":
		{
			fmt.Println("running in 32 bit mode")
		}
	case "AMD64":
		{
			fmt.Println("Running in 64 bit mode")
		}
	case "IA64" :
		{
			fmt.Println("Remember IA64?")
		}
	}
	fmt.Println("Ran our process!")
}
