package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

const (
	Help_msg string = "p:\n\tCreate new person or print info about current person\n\tFormat: p [name] [age] [sex] [haircolor]\n\npwrite:\n\tWrite person data to binary file\n\tFormat: pwrite [filename]\n\npread:\n\tRead person data from file\n\tFormat: pread [filename]\n"
	Dir_tmp  string = "tmp/"
	Hi_msg   string = "Binary reader writer\nYou can create some Person struct (name string, age uint8, sex bool, haircolor string)\nWrite it to binary file\nRead Person data from binary file\nType 'help' for more info\n"
)

func IsDirectory(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func main() {

	// Create temp directory
	if exists, _ := IsDirectory(Dir_tmp); !exists {
		if dir_err := os.Mkdir(Dir_tmp, os.ModeDir); dir_err != nil {
			log.Fatal("Could not make temp dir ", dir_err)
		}
	}

	fmt.Print(Hi_msg)

	// Read person from file
	var person Person
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
