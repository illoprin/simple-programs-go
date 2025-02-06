package main

import (
	"fmt"
)

func main() {
	var expression string = "934+23 * 70"
	result, err := EvaluateExpression(expression)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}
	fmt.Printf("Result is: %.2f", result)
}
