package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput() (string, bool) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "enter a valid input", false
	}
	input = strings.TrimSpace(input)
	if strings.ToLower(input) == "quit" {
		return "", true
	}
	return input, false
}

func PlayGame(maxAttempts int, bufferScore *[]int) {
	attempts := 0

	num := r.Intn(100) + 1

	start := time.Now()
	for {
		fmt.Println("Enter your guess : ")
		val, quit := readInput()
		if quit {
			fmt.Println("You chose to exit the game. Goodbye!")
			os.Exit(0)
		}
		valNum, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Enter a valid number")
			continue
		}

		attempts++
		if attempts > maxAttempts {
			fmt.Println("Number of attempts exhausted")
			break
		}
		if valNum == num {
			duration := time.Since(start)
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n ", attempts)
			fmt.Printf("you guessed the answer in %.2f seconds\n", duration.Seconds())
			highestScore := maxAttempts - attempts
			*bufferScore = append(*bufferScore, highestScore)
			n := (HighestScore(*bufferScore))
			fmt.Printf(" The highest score is %d\n", n)

			break
		}
		if valNum > num {
			fmt.Printf("Incorrect! The number is less than %d\n", valNum)
			fmt.Printf("You have %d more attempts\n", maxAttempts-attempts)
			Hint(num)

		}
		if valNum < num {
			fmt.Printf("Incorrect! The number is greater than %d\n", valNum)
			fmt.Printf("You have %d more attempts\n", maxAttempts-attempts)
			Hint(num)

		}

	}
	fmt.Printf("Game over! The correct number was %d\n", num)

}
func Hint(num int) {
	if num%2 == 0 && num%5 == 0 {
		fmt.Println("The number is even and divisible by 5")
	} else if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

}
func HighestScore(container []int) int {
	if len(container) == 0 {
		return 0
	}
	maxScore := container[0]
	for _, r := range container {
		if r > maxScore {
			maxScore = r
		}
	}
	return maxScore
}
