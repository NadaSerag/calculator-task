package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, readerPtr *bufio.Reader) string {
	fmt.Println(prompt)
	input, _ := readerPtr.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

// to run type in terminal (powershell): go run calc.go
func main() {

	reader := bufio.NewReader(os.Stdin)

	//var expression string

	calculatorChoice := getInput("Which calculator: (a) Two Values, (b) Whole Expression", reader)

	if calculatorChoice == "a" {

	} else {

	}
}
