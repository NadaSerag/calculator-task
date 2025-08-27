package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Rule of thumb
// Lowercase: private to the package.
// Capitalized: exported, visible outside the package.

func parseString(stringToParse string) []string {
	parsedString := strings.Split(stringToParse, " ")
	return parsedString
}

func evaluate(tokens []string) float64 {
	// Handling * and / in the first pass
	var stack []string
	var midResult, finalResult, firstNum, secondNum float64

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "*" || token == "/" {
			// taking the previous number (firstNum in calculation)
			firstNum, _ = strconv.ParseFloat(tokens[i-1], 64)
			secondNum, _ = strconv.ParseFloat(tokens[i+1], 64)
			stack = stack[:len(stack)-1] // popping last number

			if token == "*" {
				midResult = firstNum * secondNum
			} else {
				if secondNum == 0 {

					return 0
				}
				midResult = firstNum / secondNum
			}

			//returning midResult (float64) into a string rto be able to append it to the stack array which is a string
			stack = append(stack, strconv.FormatFloat(midResult, 'f', 2, 64))
			i++ // skip next number
		} else {
			stack = append(stack, token)
		}
	}

	// Handling + and - in the first pass
	finalResult, _ = strconv.ParseFloat(stack[0], 64)
	for i := 1; i < len(stack); i += 2 {
		operator := stack[i]
		num, _ := strconv.ParseFloat(stack[i+1], 64)
		if operator == "+" {
			finalResult += num
		} else if operator == "-" {
			finalResult -= num
		}
	}
	return finalResult
}

func Advanced_calc() {
	reader := bufio.NewReader(os.Stdin)

	expression := getInput("Enter an expression: ", reader)
	expTokenized := parseString(expression)
	result := evaluate(expTokenized)

	resultString := strconv.FormatFloat(result, 'f', 2, 64)

	fmt.Println(expression + " " + " = " + resultString)
}
