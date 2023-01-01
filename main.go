package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	helpFlag   = false
	helpString = `You can use this program rvdy by running the program as "rvdy <Some future date>" example: "rvdy 23 Dec 2023"`
)

func init() {
	flag.BoolVar(&helpFlag, "help", false, helpString)
}

func main() {
	flag.Parse()
	if helpFlag {
		flag.PrintDefaults()
		os.Exit(1)
	}
	args := flag.Args()
	a := strings.Join(args, " ")
	futureDate, err := time.Parse("2 Jan 2006", a)
	if err != nil {
		fmt.Println("You need to type the future date in the format dd MMM YYY. Example: 23 Jan 2024")
		os.Exit(1)
	}
	fmt.Println(time.Until(futureDate))
}
