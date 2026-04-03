package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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
	bufferScore := []int{}

	num := r.Intn(100) + 1

	start := time.Now()
	for {
		fmt.Println("Enter your guess : ")
		val, err := readInput()
		if err != nil {
			fmt.Println("Enter only valid numbers")
			continue
		}

		attempts++
		if attempts >= maxAttempts {
			fmt.Println("Number of attempts exhausted")
			break
		}
		if val == num {
			duration := time.Since(start)
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n ", attempts)
			fmt.Printf("you guessed the answer in %.2f seconds\n", duration.Seconds())
			highestScore := maxAttempts - attempts
			bufferScore = append(bufferScore, highestScore)
			n := (HighestScore(bufferScore, highestScore))
			fmt.Println(n)

			break
		}
		if val > num {
			fmt.Printf("Incorrect! The number is less than %d\n", val)
			fmt.Printf("You have %d more attempts\n", maxAttempts-attempts)
			Hint(num)

		}
		if val < num {
			fmt.Printf("Incorrect! The number is greater than %d\n", val)
			fmt.Printf("You have %d more attempts\n", maxAttempts-attempts)
			Hint(num)

		}

	}

}
func Hint(num int) {
	if num%2 == 0 && num%5 == 0 {
		fmt.Println("The number is even and divisible by 5")
	}
	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
	if num%3 == 0 {
		fmt.Println("The number is divisible by 3")
	}

}
func HighestScore(container []int, highestScore int) int {
	newHighestScore := 0
	for _, r := range container {
		if r < highestScore {
			newHighestScore = r
		}
	}
	return newHighestScore
}
