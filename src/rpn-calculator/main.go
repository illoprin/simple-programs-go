package main

import (
	"fmt"
	//	"strings"
)

func main() {
	fmt.Println("Reverse Polish Notation Based Calculator")
	fmt.Println("\nCommands:\ncalc\n\tEvaluate expression\n\tOperators: + - * / ^\n\tNO SPACES!!!\n\nexit\n\tClose program")
	for true {
		fmt.Printf("> ")

		// Read input command string
		var cmd, args string
		fmt.Scanln(&cmd, &args)

		if cmd == "exit" {
			break
		} else if cmd == "calc" {
			result, err := EvaluateExpression(args)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Result is: %.2f\n", result)
		}
	}
	fmt.Println("Bye :)")
}
