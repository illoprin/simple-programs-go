package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func EvaluateExpression(expr string) (float64, error) {
	tokens, err := getTokens(expr)
	if err == nil {
		var rpn []string = getReversePolishNotation(tokens)
		return evaluateRPN(rpn), nil
	}
	return 0.0, err
}

func evaluateRPN(rpn []string) float64 {
	stack := New()

	for _, token := range rpn {

		// Process expression if token is operator
		// Add add result to stack
		if isOperator(token) {
			second, _ := stack.Pop().(float64)
			first, _ := stack.Pop().(float64)
			ans, err := eval(first, second, token)
			if err == nil {
				stack.Push(ans)
			} else {
				fmt.Println(err)
			}
			continue
		}

		// Else - add number to stack
		num, _ := strconv.ParseFloat(token, 64)
		stack.Push(float64(num))
	}
	return stack.Pop().(float64)
}

func getTokens(expr string) ([]string, error) {
	var current strings.Builder
	tokens_list := make([]string, 0, 10)

	// Iterate through string by characters
	var err error = nil
	for i, char := range expr {
		// Write separate number to tokens list
		if unicode.IsDigit(char) || string(char) == "." {
			current.WriteRune(char)
			continue
		} else if current.Len() > 0 {
			tokens_list = append(tokens_list, current.String())
			current.Reset()
		}

		// Add operator to tokens list
		if getTokenWeight(string(char)) != 0 {
			tokens_list = append(tokens_list, string(char))
		} else if string(char) == " " {
			// Skip spaces
			continue
		} else {
			// Unexpected token
			err = fmt.Errorf("Unexpected token at %d: %s\n", i, string(char))
		}
	}

	if current.Len() > 0 {
		tokens_list = append(tokens_list, current.String())
		current.Reset()
	}

	return tokens_list, err
}

func getReversePolishNotation(tokens []string) []string {
	var rpn []string = make([]string, 0, 10)

	operators := New() // Operands stack

	var priority int8
	for _, token := range tokens {
		priority = getTokenWeight(token)
		if priority == 0 {
			rpn = append(rpn, token)
		} else if priority == 1 {
			operators.Push(token)
		} else if priority > 1 {

			// Add operator to stack
			// Idea:
			// 	if priority of last operator in stack is less then current token priority
			// 		we push it to stack
			//  else
			//  	we pop elements until we encounter an operator with lower priority
			// 		while simultaneously adding operators to the final RPN expression

			for operators.Len() > 0 {
				if getTokenWeight(operators.Peek().(string)) >= priority {
					rpn = append(rpn, operators.Pop().(string))
				} else {
					break
				}
			}
			operators.Push(token)
		} else if priority == -1 {

			// closing bracket case
			// Idea:
			// 		find opening bracket in stack and remove it
			// 		while simultaneously adding operands to the final RPN expression
			for operators.Len() > 0 &&
				getTokenWeight(operators.Peek().(string)) != 1 {

				rpn = append(rpn, operators.Pop().(string))
			}
			operators.Pop()
		}
	}

	// Add remaining operators to RPN expression
	for operators.Len() != 0 {
		rpn = append(rpn, operators.Pop().(string))
	}

	return rpn
}

func isOperator(token string) bool {
	return getTokenWeight(token) != 0
}

func eval(o1 float64, o2 float64, o string) (float64, error) {
	switch o {
	case "+":
		return o1 + o2, nil
	case "-":
		return o1 - o2, nil
	case "*":
		return o1 * o2, nil
	case "/":
		return o1 / o2, nil
	case "^":
		return math.Pow(o1, o2), nil
	}
	return 0.0, fmt.Errorf("Unexpected operator %s\n", o)
}

func getTokenWeight(token string) int8 {
	if token == "^" {
		return 4
	} else if token == "*" || token == "/" {
		return 3
	} else if token == "+" || token == "-" {
		return 2
	} else if token == "(" {
		return 1
	} else if token == ")" {
		return -1
	} else {
		// 0-9 numbers case
		return 0
	}
}
