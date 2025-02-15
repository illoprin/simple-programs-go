package main

import "fmt"

type llist struct {
	length uint64
	begin  *llist_node
	end    *llist_node
}

type llist_node struct {
	value interface{}
	next  *llist_node
}

// === Getters
func (this *llist) Len() uint64 {
	return this.length
}
func (this *llist) Begin() *llist_node {
	return this.begin
}
func (this *llist) End() *llist_node {
	return this.end
}

func NewLlist() *llist {
	return &llist{
		0, nil, nil,
	}
}

// Push new element at the end of list
func (this *llist) PushBack(value interface{}) {
	new_node := llist_node{
		value: value, next: nil,
	}

	if this.length >= 1 {
		this.end.next = &new_node
		this.end = &new_node
	} else if this.length == 0 {
		this.end = &new_node
		this.begin = &new_node
	}

	this.length++
}

// Push new element at the begin of list
func (this *llist) PushFront(value interface{}) {
	new_node := llist_node{
		value: value, next: nil,
	}

	if this.length == 0 {
		this.begin = &new_node
		this.end = &new_node
	} else if this.length > 0 {
		new_node.next = this.begin
		this.begin = &new_node
	}

	this.length++
}

// Push new value after node with index
func (this *llist) PushAfter(index uint64, value interface{}) {
	// Process simple cases
	if this.length == 0 && index > 0 {
		return
	} else if index == 0 {
		this.PushFront(value)
		return
	} else if index == this.length-1 {
		this.PushBack(value)
		return
	} else if index > this.length {
		return
	}

	new_node := llist_node{
		value: value, next: nil,
	}

	var counter uint64 = 0
	node := this.begin
	for counter != index {
		node = node.next
		counter++
	}

	node_next := node.next
	node.next = &new_node
	new_node.next = node_next

	this.length++
}

// Delete last element in list and return this value
func (this *llist) PopBack() interface{} {
	if this.length == 0 {
		return nil
	}

	value_to_delete := this.end.value
	// Process simple cases (when len of list is less then 3)
	if this.length == 1 {
		this.begin = nil
		this.end = nil
	} else if this.length == 2 {
		this.end = this.begin
		this.begin.next = nil
	} else if this.length > 2 {
		// Find previous last element
		var previous_last *llist_node = this.begin
		for previous_last.next.next != nil {
			previous_last = previous_last.next
		}

		// Update pointers
		this.end = previous_last
		previous_last.next = nil
	}

	this.length--

	return value_to_delete
}

func (this *llist) PopFront() interface{} {
	if this.length == 0 {
		return nil
	}
	value_to_delete := this.begin.value

	if this.length == 1 {
		this.begin = nil
		this.end = nil
	} else if this.length > 1 {
		this.begin = this.begin.next
	}

	this.length--

	return value_to_delete
}

func (this *llist) Print() {
	node := this.begin
	fmt.Print("{ ")
	for node != nil {
		fmt.Print(node.value, " ")
		node = node.next
	}
	fmt.Print("}\n")
}
