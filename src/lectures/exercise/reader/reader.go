//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

// This is my solution:
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	linesNum, commandsNum := 0, 0

	for {
		input, inputErr := r.ReadString('\n')
		s := strings.TrimSpace(input)
		if s == "" {
			linesNum += 1
			continue
		}
		if s == "hello" {
			fmt.Println("Hello, User! What's up?")
			linesNum += 1
			commandsNum += 1
		} else if s == "bye" {
			fmt.Println("Bye, User!")
			linesNum += 1
			commandsNum += 1
		} else if s == "Q" || s == "q" {
			break
		}

		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
		}
	}
	fmt.Println("The number of non-blank lines entered:", linesNum)
	fmt.Println("The number of commands entered:", commandsNum)
}

//This is professors solution:
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// const (
// 	CmdHello = "hello"
// 	CmdGoodbye = "bye"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	numLines, numCommands := 0, 0

// 	for scanner.Scan() {
// 		if strings.ToUpper(scanner.Text()) == "Q" {
// 			break
// 		} else {
// 			text := strings.TrimSpace(scanner.Text())
// 			switch text {
// 			case CmdHello:
// 				numCommands += 1
// 				fmt.Println("command response: hi")
// 			case CmdGoodbye:
// 				numCommands += 1
// 				fmt.Println("command response: hi")
// 			}
// 			if text != "" {
// 				numLines += 1
// 			}
// 		}
// 	}
// 	fmt.Printf("You entered %v lines\n", numLines)
// 	fmt.Printf("You entered %v commands\n", numCommands)
// }