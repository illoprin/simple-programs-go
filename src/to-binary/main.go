package main

import (
	"fmt"
)

const Help_msg string = "p:\n\tCreate new person or print info about current person\n\tFormat: p [name] [age] [sex] [haircolor]\n\npwrite:\n\tWrite person data to binary file\n\tFormat: pwrite [filename]\n\npread:\n\tRead person data from file\n\tFormat: pread [filename]\n"

func main() {

	// Read person from file
	var person Person
	_ = person
	for true {
		fmt.Printf("> ")

		// Read input command string
		var cmd string
		fmt.Scan(&cmd)

		if cmd == "exit" {
			break
		} else if cmd == "p" {
			// Create new person or show info about person
			_, err := fmt.Scanln(&person.Name, &person.Age, &person.Sex, &person.HairColor)

			if err != nil {
				person.Print()
				continue
			}

		} else if cmd == "pwrite" {

			// Write current person to file
			var filename string
			fmt.Scanln(&filename)
			person.WriteToBIN(filename)

		} else if cmd == "pread" {

			// Read file to current person
			var filename string
			fmt.Scanln(&filename)
			person.ReadFromBIN(filename)

		} else if cmd == "pclear" {
			person = Person{}
		} else if cmd == "help" {
			// Show help message
			fmt.Print(Help_msg)
		}
	}
}
