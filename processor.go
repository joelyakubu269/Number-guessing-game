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
func Easy(n, num int) {

	if n == num {
		fmt.Println("Congratulations! You guessed the correct number .")
	}
	if n > num {
		fmt.Printf("Incorrect! The number is less than %d\n", n)
	}
	if n < num {
		fmt.Printf("Incorrect! The number is greater than %d\n", n)
	}
}
