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

func Basic_2values_calc() {
	reader := bufio.NewReader(os.Stdin)

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
}
