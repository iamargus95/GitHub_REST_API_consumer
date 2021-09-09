package main

import (
	"flag"
)

func main() {

	var con bool
	flag.BoolVar(&con, "con", false, "Runs the application concurrently.")

	flag.Parse()
	usernames := flag.Args()

	if con {
		Concurrently(usernames)
	} else {
		Sequence(usernames)
	}

}
