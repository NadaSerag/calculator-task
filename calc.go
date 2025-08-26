package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, readerPtr *bufio.Reader) string {
	fmt.Println(prompt)
	input, _ := readerPtr.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

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

// to run type in terminal (powershell): go run calc.go
func main() {

	//var expression string
	reader := bufio.NewReader(os.Stdin)

	calculatorChoice := getInput("Which calculator: (a) Two Values, (b) Whole Expression", reader)

	if calculatorChoice == "a" {
		var firstNum, secondNum, result float64
		var operator, input, input2, resultString, concatenatedString string
		var divByZero, historyUsedBefore bool

		historyUsedBefore = false

		//creating "history" slice
		var history = []string{}

		divByZero = false

		for {
			//(a) Calculate, (b) View History, (c) Clear History, (d) Exit
			option := getInput("Choose what you want to do: (a) Calculate, (b) View History, (c) Clear History, (d) Exit", reader)

			switch option {
			case "a":
				{
					fmt.Print("Enter first number: , To return to the menu, Enter 'm' ")

					input, _ = reader.ReadString('\n')
					input = strings.TrimSpace(input) // removes newline
					if input == "m" {
						break
					}

					firstNum, _ = strconv.ParseFloat(input, 64)

					fmt.Print("Enter operator (+ - * /): ")
					operator, _ = reader.ReadString('\n')
					operator = strings.TrimSpace(operator)

					for {
						if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
							fmt.Println("Unknown operator:", operator)
							fmt.Print("Enter a VALID operator (+ - * /): ")
							operator, _ = reader.ReadString('\n')
							operator = strings.TrimSpace(operator)
						} else {
							break
						}
						// fmt.Print("Enter a VALID operator (+ - * /): ")
						// fmt.Scan(&operator)
					}

					fmt.Print("Enter second number: ")
					input2, _ = reader.ReadString('\n')
					input2 = strings.TrimSpace(input2)
					secondNum, _ = strconv.ParseFloat(input2, 64)
					//fmt.Scan(&secondNum)

					//reader := bufio.NewReader(os.Stdin)
					//fmt.Print("Enter what you want to calculate: ")

					//	fmt.Print(reader) <---- this print all this:
					// Enter what you want to calculate: &{[0 0 0 0 0 0 0 0
					//  0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
					//  ... infinite zeros in between ... 0000] 0xc000054020 0 0 <nil> -1 -1}

					//That why we have to read what was read again
					// and tell it to stop reading after end line \n character:

					//The underscore: to ignore the error return value rom the ReadString function.
					// the function ReadString('\n') actually returns two values:
					// string → the data read up to and including the delimiter ('\n' here).
					// error → an error value (which is nil if everything went fine).

					// readerOfreader, _ := reader.ReadString('\n')
					// fmt.Print(readerOfreader)

					switch operator {
					case "+":
						result = firstNum + secondNum
					case "-":
						result = firstNum - secondNum
					case "*":
						result = firstNum * secondNum
					case "/":
						if secondNum == 0 {
							divByZero = true
							fmt.Println("Error: Can't divide by zero!")
							break
						}
						result = firstNum / secondNum
						// default:
						// 	fmt.Println("Unknown operator:", operator)
						// 	return
					}

					if !divByZero {
						fmt.Println("Result: ", result)
					} else {
						fmt.Println("didnt enter if")
					}

					resultString = strconv.FormatFloat(result, 'f', 2, 64)
					concatenatedString = input + " " + operator + " " + input2 + " = " + resultString
					history = append(history, concatenatedString)
					divByZero = false
					historyUsedBefore = true
				}

			case "b":
				{
					fmt.Println(history)
				}

			case "c":
				{
					history = []string{}
					if !historyUsedBefore {
						fmt.Println("History double checked to be cleared (was already empty)")
					} else {
						fmt.Println("History cleared successfully")
					}
				}

			case "d":
				{
					fmt.Println("Exited Program")
					return
				}

			default:
				{
					fmt.Println("Invalid menu option chosen, Lets try again:")
				}

			}
		}
	} else {
		expression := getInput("Enter an expression: ", reader)
		expTokenized := parseString(expression)
		result := evaluate(expTokenized)

		resultString := strconv.FormatFloat(result, 'f', 2, 64)

		fmt.Println(expression + " " + " = " + resultString)

	}
}
