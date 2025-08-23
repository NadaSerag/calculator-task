package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// to run type in terminal (powershell): go run calc.go
func main() {

	reader := bufio.NewReader(os.Stdin)
	var firstNum, secondNum, result float64
	var operator, input, input2 string
	var divByZero bool

	divByZero = false

	for {
		fmt.Print("Enter first number: , To exit, Enter 'e' ")

		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input) // removes newline
		if input == "e" {
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

		divByZero = false
	}
}
