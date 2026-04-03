package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println(" You have 5 chances to guess the correct number.")
	fmt.Println("In order to quit enter 4")

	fmt.Println(`Please select the difficulty level:
		1. Easy (10 chances)
		2. Medium (5 chances)
		3. Hard (3 chances)
		`)
	bufferScore := []int{}
	for {
		fmt.Println("Enter your choice: ")
		input, err := readInput()
		if err {
			fmt.Println("Goodbye")
			return

		}
		switch input {
		case "1":

			fmt.Println(`Great! You have selected the Easy difficulty level.
				Let's start the game!`)
			PlayGame(10, &bufferScore)

		case "2":
			fmt.Println(`Great! You have selected the Medium difficulty level.
				Let's start the game!`)
			PlayGame(5, &bufferScore)

		case "3":
			fmt.Println(`Great! You have selected the Hard difficulty level.
				Let's start the game!`)
			PlayGame(3, &bufferScore)
		case "quit":
			return

		default:
			fmt.Println("Unknown choice enter a valid selection")
		}
		fmt.Println("Do you wish to continue (yes/no): ")
		input, val := reader.ReadString('\n')
		if val != nil {
			fmt.Println("Enter a valid input")
			continue
		}
		input1 := strings.TrimSpace(input)
		if strings.ToLower(input1) != "yes" {
			fmt.Println("Goodbye")
			break
		}
	}
}
