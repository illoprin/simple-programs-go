package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var list *Llist = NewLlist()
	var items1 []int = []int{2, 3, 4, 5, 3, 10, 23, 222, 10}
	var items2 []int = []int{34, 32, 45}
	var items3 []int = []int{84, 82, 81}

	for i := 0; i < len(items1); i++ {
		list.PushBack(items1[i])
	}
	list.Print()
	for i := 0; i < len(items2); i++ {
		list.PushFront(items2[i])
	}
	list.Print()
	for i := 0; i < len(items3); i++ {
		list.PushAfter(list.Len()/2, items3[i])
	}
	list.Print()

	fmt.Println()

	for i := 0; i < len(items1)/2; i++ {
		fmt.Printf("Deleting items: %d\n", list.PopBack().(int))
		fmt.Printf("Len: %d\n", list.Len())
	}

	for i := 0; i < len(items1)/2; i++ {
		fmt.Printf("Deleting items: %d\n", list.PopFront().(int))
		fmt.Printf("Len: %d\n", list.Len())
	}

	for i := 0; i < 20; i++ {
		list.Print()
		if rand.Intn(2) == 1 {
			fmt.Println("Deleting items: ", list.PopFront())
		} else {
			fmt.Println("Deleting items: ", list.PopBack())
		}
	}

}
