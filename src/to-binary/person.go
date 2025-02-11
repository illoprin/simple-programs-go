package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Person struct {
	Name      string
	Age       uint8
	Sex       bool
	HairColor string
}

func (this *Person) WriteToBIN(filename string) {
	file, err := os.Create("tmp/" + filename + ".pbin")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	defer file.Close()

	// Write name string len
	var name_len uint8 = uint8(len(this.Name))
	binary.Write(file, binary.LittleEndian, name_len)

	// Write name string
	var name []byte = []byte(this.Name)
	binary.Write(file, binary.LittleEndian, name)

	// Write age
	binary.Write(file, binary.LittleEndian, this.Age)

	// Write sex bool
	binary.Write(file, binary.LittleEndian, this.Sex)

	// Write hair color string len
	var hair_color_len uint8 = uint8(len(this.HairColor))
	binary.Write(file, binary.LittleEndian, hair_color_len)

	// Write hair color string
	var hair_color []byte = []byte(this.HairColor)
	binary.Write(file, binary.LittleEndian, hair_color)
}

func (this *Person) ReadFromBIN(filename string) {
	file, err := os.Open("tmp/" + filename + ".pbin")
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}
	defer file.Close()

	// Read name (string)
	var name_str_len uint8 // Read string len
	if err := binary.Read(file, binary.LittleEndian, &name_str_len); err == nil {
		name_str_bytes := make([]byte, name_str_len) // Read string bytes
		binary.Read(file, binary.LittleEndian, name_str_bytes)

		this.Name = string(name_str_bytes)
	} else {
		fmt.Println("Error while reading file: ", err)
		return
	}

	// Read age (uint8)
	binary.Read(file, binary.LittleEndian, &this.Age)

	// Read sex (bool)
	binary.Read(file, binary.LittleEndian, &this.Sex)

	// Read hair color (string)
	var hair_color_str_len uint8
	if err := binary.Read(file, binary.LittleEndian, &hair_color_str_len); err == nil {
		hair_color_str_bytes := make([]byte, hair_color_str_len)
		binary.Read(file, binary.LittleEndian, hair_color_str_bytes)

		this.HairColor = string(hair_color_str_bytes)
	} else {
		fmt.Println("Error while reading file: ", err)
		return
	}
}

func (this *Person) Print() {
	fmt.Printf("Person {\n\tName: %s\n\tAge: %d\n\tSex: %v\n\tHair Color: %s\n}\n",
		this.Name, this.Age, this.Sex, this.HairColor)
}
