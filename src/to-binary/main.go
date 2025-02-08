package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Read person from file

	var person Person
	_ = person
	for true {
		fmt.Printf("> ")

		// Read input command string
		var cmd, args string
		fmt.Scanln(&cmd, &args)

		if cmd == "exit" {
			break
		} else if cmd == "person" {
			// Create new person or show info about person

		} else if cmd == "person_write" {
			// Write current person to file

		} else if cmd == "person_read" {
			// Read file to current person
		}
	}
}
