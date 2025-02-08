package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"strconv"
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

const (
	Dir_tmp  string = "tmp"
	Help_msg string = "gen:\n\tGenerate image and write it to tmp directory\n\teg: gen [width] [height] [type]\n\tImage types:\n\t\tbox\n\t\tnoise\n\t\tgrass\n\t\trock\nexit:\n\tClose program\nhelp:\n\tShow help message"
	Gen_inc  string = "Incorrect input\nFormat:\n\tgen [width] [height] [type]"
)

func main() {
	// Create temp directory
	if exists, _ := IsDirectory(Dir_tmp); !exists {
		if dir_err := os.Mkdir(Dir_tmp, os.ModeDir); dir_err != nil {
			log.Fatal("Could not make temp dir ", dir_err)
		}
	}

	for {
		// Process user input
		fmt.Print("> ")
		var cmd string
		fmt.Scan(&cmd)

		if cmd == "exit" {
			break
		} else if cmd == "help" {
			fmt.Print(Help_msg)
		} else if cmd == "gen" {
			var (
				img_width, img_height uint16
				image_type            string
			)
			if n, _ := fmt.Scan(&img_width, &img_height, &image_type); n != 3 {
				fmt.Println(Gen_inc)
				continue
			}
			img_obj, err := CreateTexture(img_width, img_height, image_type)
			if err != nil {
				fmt.Println(err)
				continue
			}
			var filepath string = "tmp/" + strconv.Itoa(rand.Intn(1000)) + ".png"
			if file_err := WriteImageToPNG(img_obj, filepath); file_err != nil {
				fmt.Println(file_err)
				continue
			}
			fmt.Println("Image writed to path ", filepath)
		}
	}
	fmt.Println("Bye :)")
}
