package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Sence: we have some text files with numbers to read. Necessary to calculate sum of numbers for each file

const (
	Hi_msg string = "Parallel file generator and reader\n'help' for more info"

	Help_msg string = "fcreate:\n\tGenerate some files with numbers separated with space\n\tFormat: fcreate [...filename]\nfsum:\n\tCalculate sum of numbers for each file\n\tFormat: fsum [...filename]\nmrand:\n\tDefine max values for numbers in generated files\n\tFormat: mrand [max_nums int] [max_value int]\nexit:\n\tClose program"

	Err_msg string = "Unknown command, type 'help' for reference"
)

func main() {
	fmt.Println(Hi_msg)

	stdin_reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("> ")
		input, _ := stdin_reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		args := parts[1:]

		switch command {
		case "exit":
			fmt.Println("Bye :)")
			return
		case "help":
			fmt.Println(Help_msg)
		case "fcreate":
			if len(args) == 0 {
				fmt.Println("Command 'fcreate' takes some arguments")
				continue
			}

			bytesize, err := WriteNumbersToFiles(args)
			if err != nil {
				fmt.Printf("Error %v\n", err)
			}

			fmt.Printf("Writed %d bytes\n", bytesize)
		case "fsum":
			if len(args) == 0 {
				fmt.Println("Command 'fcreate' takes some arguments")
				continue
			}

			// For each file we read string slice and calculate sum of ints from slice
			// after that we add this number to whole sum
			var sum int64 = 0
			for _, filename := range args {
				nums, ferr := ReadNumbersFromFile(filename)
				if ferr != nil {
					fmt.Printf("Error %v\n", ferr)
					continue
				}
				sum += SumFromSlice(nums)
			}
			fmt.Printf("Sum of whole numbers: %d\n", sum)

		case "mrand":
			val, err := strconv.Atoi(args[0])
			if err == nil {
				Max_rand_value = uint32(val)
			}
			val, err = strconv.Atoi(args[1])
			if err == nil {
				Max_nums_in_file = uint16(val)
			}
		default:
			fmt.Println(Err_msg)
		}
	}
}
