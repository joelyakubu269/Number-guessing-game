// Welcome to the Number Guessing Game!
// I'm thinking of a number between 1 and 100.
// You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

// Enter your choice: 2

// Great! You have selected the Medium difficulty level.
// Let's start the game!

// Enter your guess: 50
// Incorrect! The number is less than 50.

// Enter your guess: 25
// Incorrect! The number is greater than 25.

// Enter your guess: 35
// Incorrect! The number is less than 35.

// Enter your guess: 30
// Congratulations! You guessed the correct number in 4 attempts.
// To make the game more interesting, you can add the following features:

//     Allow the user to play multiple rounds of the game (i.e., keep playing until the user decides to quit). You can do this by asking the user if they want to play again after each round.

//     Add a timer to see how long it takes the user to guess the number.

//     Implement a hint system that provides clues to the user if they are stuck.

//     Keep track of the user's high score (i.e., the fewest number of attempts it took to guess the number under a specific difficulty level).

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"unicode"
	"math/rand"
	"time"
	"strconv"

)
func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println(" You have 5 chances to guess the correct number.")
	reader:= bufio.NewReader(os.Stdin)
	 rand.Seed(time.Now().UnixNano()) // this has no return value
	num := rand.Intn(100)  + 1
	for {
		fmt.Println(`Please select the difficulty level:
		1. Easy (10 chances)
		2. Medium (5 chances)
		3. Hard (3 chances)
		`)
	    }
		fmt.Println("Enter your choice: ")
		readInput()
		notAlpha:= true
		for _,r:= range input {
			if !unicode.IsDigit(r) {
				notAlpha= false
				break
			}
			if !notAlpha {
				fmt.Println("Enter only digits")
				continue
			}
		}
		switch input {
		case "1" :
			fmt.Println(`Great! You have selected the Easy difficulty level.
				Let's start the game!`)
				attempt:= 0
			fmt.Println("Enter your guess : ")
				val := readInput()
				attempt++
				Easy(val)

			
		case "2" :
			fmt.Println(`Great! You have selected the Medium difficulty level.
				Let's start the game!`)
		case "3" :
			fmt.Println(`Great! You have selected the Hard difficulty level.
				Let's start the game!`)
		default:
			fmt.Println("Unknown choice enter a valid selection")
		}
         }
func readInput() string{
	input,err := reader.ReadString('\n')
		if err!= nil {
			fmt.Println("Error:",err)
		}
		input= strings.TrimSpace(input)
		return input
}
func Easy(input string) {
	n,err:= strconv.Atoi(input)
	if err!= nil {
		fmt.Println("Error:",err)
	}
	if n == num {
		fmt.Println("Congratulations! You guessed the correct number .")
	}
	if n > num {
	fmt.Printf("Incorrect! The number is less than %d\n",n)
	}
	if n < num {
		fmt.Printf("Incorrect! The number is greater than %d\n",n)
	}
}