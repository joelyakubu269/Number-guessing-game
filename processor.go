package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readInput() (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	}
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// func Easy(n, num, attempt int) {

// 	if n == num {
// 		fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n .", attempt)

//		}
//		if n > num {
//			fmt.Printf("Incorrect! The number is less than %d\n", n)
//			fmt.Printf("You have %d more attempts\n", 10-attempt)
//		}
//		if n < num {
//			fmt.Printf("Incorrect! The number is greater than %d\n", n)
//			fmt.Printf("You have %d more attempts\n", 10-attempt)
//		}
//	}

// ✅ Replace Easy/Medium/Hard with ONE function
// ✅ Create playGame(maxAttempts)
// ✅ Move num inside the game

func PlayGame(maxAttempts int) {
	attempts := 0
	for {
		fmt.Println("Enter your guess : ")
		val, err := readInput()
		if err != nil {
			fmt.Println("Enter only valid numbers")
			continue
		}
		attempts++
		if attempts == maxAttempts {
			fmt.Println("Number of attempts exhausted")
			break
		}
		if val == num {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n .", attempts)

			break
		}
		if val > num {
			fmt.Printf("Incorrect! The number is less than %d\n", val)
			fmt.Printf("You have %d more attempts\n", maxAttempts-attempts)
		}
		if val < num {
			fmt.Printf("Incorrect! The number is greater than %d\n", val)
			fmt.Printf("You have %d more attempts\n", maxAttempts-attempts)
		}

	}

}
