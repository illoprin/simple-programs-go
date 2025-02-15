package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	Max_rand_value   uint32 = 64000
	Max_nums_in_file uint16 = 100
)

// Returns error while writing and writed bytes
func WriteNumbersToFiles(filenames []string) (uint64, error) {
	var bsum uint64 = 0 // =Created files byte size
	for _, filename := range filenames {
		// Create file
		file, err := os.Create("tmp/" + filename + ".txt")
		if err != nil {
			return 0, err
		}

		// Write numbers
		var max int = int(Max_nums_in_file)/3 + rand.Intn(int(Max_nums_in_file)-int(Max_nums_in_file)/3)
		for i := 0; i < max; i++ {
			fmt.Fprint(file, strconv.Itoa(rand.Intn(int(Max_rand_value)))+" ")
		}

		// Get file size
		fs, _ := file.Stat()
		bsum += uint64(fs.Size())
		file.Close()
	}
	return bsum, nil
}

func ReadNumbersFromFile(filename string) ([]int, error) {
	fbytes, err := os.ReadFile("tmp/" + filename + ".txt")
	if err != nil {
		return nil, err
	}

	// Read string and convert it to space separated array
	str := string(fbytes)
	tokens := strings.Fields(str)

	// Create int slice from string array
	var nums []int = make([]int, len(tokens))
	var num int

	// Set values of ints from converting string array
	for i := 0; i < len(nums); i++ {
		num, err = strconv.Atoi(tokens[i])

		if err != nil {
			return nil, err
		}

		nums[i] = num
	}

	return nums, nil
}
