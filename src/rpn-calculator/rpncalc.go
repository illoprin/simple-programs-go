package main

import (
	"fmt"
	"strings"
	"unicode"
)

func EvaluateExpression(expr string) (float32, error) {
	var tokens []string = getTokensList(expr)
	var rpn []string = getReversePolishNotation(tokens)
	return 0.0, nil
}

func getReversePolishNotation(tokens []string) []string {
	var rpn []string = make([]string, 0, 10)

	var number_tokens []string = make([]string, 0, 10)
	operands := New() // Operands stack

	for _, token := range tokens {
		if !isOperand(token) {
			number_tokens = append(number_tokens, token)
		} else {
			// Check priority of previous operands

			operands.Push(token)
		}
	}

	return rpn
}

func getTokensList(expr string) []string {
	var current strings.Builder
	tokens_list := make([]string, 0, 10)
	// Iterate through string by characters
	for i, char := range expr {
		if unicode.IsDigit(char) {
			current.WriteRune(char)
			continue
		}
		current.Reset()
		if isOperand(string(char)) {
			tokens_list = append(tokens_list, current.String())
		} else {
			fmt.Printf("Unexpected token at %d: %c\n", i, char)
		}
	}
	return tokens_list
}

func isOperand(token string) bool {
	if token == "*" ||
		token == "/" ||
		token == "-" ||
		token == "+" ||
		token == "(" ||
		token == ")" ||
		token == "^" {
		return true
	}
	return false
}

func getTokenWeight(r_token rune) int8 {
	token := string(r_token)
	if token == "*" || token == "/" {
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
