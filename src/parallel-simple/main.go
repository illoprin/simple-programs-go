package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const N int = 20_000

func lost_update_example() {
	wg := sync.WaitGroup{}
	wg.Add(N)

	var value int = 0
	for i := 0; i < N; i++ {
		// Goroutines overlap each other during the update process
		// Ultimately, the value will be any number up to 1000
		go func() {
			defer wg.Done()
			value++
		}()
	}

	wg.Wait()

	fmt.Printf("Done, value is %d\n", value)
}

func mutex_increment_example() {

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(N)

	var value int = 0
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()

			// Critical section
			mu.Lock()
			value++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Printf("Done, value is %d\n", value)
}

type Obj struct {
	val uint64
	str strings.Builder
}

func mutex_parallel_array_processing() {
	defer func(start time.Time) {
		fmt.Printf("Time spent: %v\n", time.Since(start))
	}(time.Now())

	fmt.Println("Parallel processing")

	var arr []Obj = make([]Obj, N)

	wg := sync.WaitGroup{}

	pull := 4
	jobs := make(chan *Obj, N)
	wg.Add(N)

	for i := 0; i < pull; i++ {
		go func() {
			for obj := range jobs {
				obj.val++
				obj.val *= 2
				obj.str.WriteString("Hello ")
				obj.str.WriteRune('😁')
				wg.Done()
			}
		}()
	}

	for i := 0; i < N; i++ {
		jobs <- &arr[i]
	}
	close(jobs)

	wg.Wait()

	fmt.Println("Done!")
}

func unparallel_array_processing() {
	defer func(start time.Time) {
		fmt.Printf("Time spent: %v\n", time.Since(start))
	}(time.Now())

	fmt.Println("Сonsistently processing")

	var arr []Obj = make([]Obj, N)

	for i := 0; i < N; i++ {
		arr[i].val++
		arr[i].val *= 2
		arr[i].str.WriteString("Hello ")
		arr[i].str.WriteRune('😁')
	}

	fmt.Println("Done!")
}

func main() {
	// Uncorrect increment example
	fmt.Println("Uncorrect increment example (using WaitGroup only)")
	lost_update_example()

	// Parallel increment example (using mutex)
	fmt.Println("Parallel increment example (using Mutex)")
	mutex_increment_example()

	unparallel_array_processing()
	mutex_parallel_array_processing()
}
