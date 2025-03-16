package main

import (
	"crypto/sha256"
	"fmt"
	"sync"
	"time"
)

func CompareHashSHA256(a [32]byte, b [32]byte) bool {
	for i := 0; i < 32; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func EnumerateOneCaseSHA256(
	hash [32]byte, length int, result chan<- []byte, wg *sync.WaitGroup) {

	defer func() {
		fmt.Printf("Case processed for length %d\n", length)
		wg.Done()
	}()

	for i := 0; i < 1<<(length*8); i++ {
		arr := make([]byte, length)
		for j := 0; j < length; j++ {
			arr[j] = byte(i >> (j * 8))
		}

		arr_hash := sha256.Sum256(arr)

		if CompareHashSHA256(arr_hash, hash) {
			result <- arr
			return
		}
	}
}

func getCollisionSHA256(hash [32]byte, maxLength int) ([]byte, error) {

	result := make(chan []byte)
	wg := sync.WaitGroup{}

	// Start gourutines for enumeration all possible cases
	for i := 0; i < maxLength; i++ {
		wg.Add(1)
		go EnumerateOneCaseSHA256(hash, i, result, &wg)
	}

	// If channel not locked after enumeration - lock it forcibly
	go func() {
		wg.Wait()
		close(result)
	}()

	// Waiting for channel result...
	select {
	// if there was a write to the channel...
	case collision := <-result:
		// return result and exit from function
		return collision, nil
	// if the result could not be found within an hour -> return error
	case <-time.After(time.Hour):
		return nil, fmt.Errorf("timeout expired")
	}
}

func main() {
	str := "hello"
	hash := sha256.Sum256([]byte(str))

	collision, err := getCollisionSHA256(hash, 6)
	if err == nil {
		fmt.Println(collision, string(collision))
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}
