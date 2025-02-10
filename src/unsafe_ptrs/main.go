package main

import (
	"fmt"
	"math/rand"
	"unsafe"
)

func main() {

	// Float32 experement
	var snums_1 []float32 = make([]float32, 5) // Float32 - 4 bytes
	for i := 0; i < 5; i++ {
		snums_1[i] = rand.Float32()
	}
	fmt.Println("Float32 Slice:", snums_1)

	var f32_byte_size uintptr = unsafe.Sizeof(snums_1[0])
	snums_b := unsafe.Pointer(&snums_1[0]) // First element of slice pointer
	for i := uintptr(0); i < uintptr(len(snums_1))*f32_byte_size; i++ {
		index := i / f32_byte_size                             // Index of element (float32)
		byte_index := i % f32_byte_size                        // Index of byte
		v := *(*float32)(unsafe.Pointer(uintptr(snums_b) + i)) // Value of float
		b := *(*byte)(unsafe.Pointer(uintptr(snums_b) + i))    // Value of byte
		if i%f32_byte_size == 0 {
			fmt.Printf("\nIndex: %d Value: float32 %.4f\n", index, v)
		}
		fmt.Printf("Index: %d Byte: %d Value: %d\n", index, byte_index, b)
	}
	fmt.Println()

	// Byte experiment
	var snums_2 []byte = make([]byte, 10)
	for i := 0; i < 10; i++ {
		snums_2[i] = byte(rand.Intn(255))
	}
	fmt.Println("Byte Slice:", snums_2)

	// sizeof(byte) - 1
	snums_2_b := unsafe.Pointer(&snums_2[0])
	for i := uintptr(0); i < uintptr(len(snums_2)); i++ {
		ptr := uintptr(snums_2_b) + i
		v := *(*byte)(unsafe.Pointer(ptr))
		fmt.Printf("Index: %d Ptr: %X Value: %d\n", i, ptr, v)
	}
	fmt.Println()

	// String experiment
	var str_byte []byte = []byte("Hello mother fucker!")
	// Create copy of byte array and cast it to string
	// var str_glyps string = string(str_byte)
	str_glyps := (*string)(unsafe.Pointer(&str_byte))
	fmt.Println("Bytes:", str_byte)
	fmt.Println("String:", *str_glyps)
	for _, char := range *str_glyps {
		fmt.Printf("Code: %d Char: %c\n", char, char)
	}
	fmt.Println()

}
