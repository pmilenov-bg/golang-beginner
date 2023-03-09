package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func GetGreeting() (greeting string)
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("Now you are going to forrest")
// 		fmt.Println("Good afternoon.")
// 	default:
// 		fmt.Println("Good evening.")
// }

func GetUserName() (name string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello, what's your name")
	name, _ = reader.ReadString('\n')
	return name
}

func getInput(message string) (input string) {
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

type User struct {
	name  string
	live  int
	money int
	// state
	//
}

// func getMessage(message-match [4]string) (output string) {
// }

func main() {
	name := GetUserName()
	fmt.Println("Good day", name)

	message := `You are on a mealow beween the town and the forrest.\nWhich way you wonna go
Press 1 for the town
Press 2 for the forrest`

	input := getInput(message)
	for true {
		switch {
		case 1:
			input == 1
			fmt.Println("Now you are going to town")
		case 2:
			input == 2
			fmt.Println("Now you are going to forrest")
		case 3:
			input == 3
		case 4:
			input == 4
		default:
			fmt.Println("wrong input")
		}

	}
}
