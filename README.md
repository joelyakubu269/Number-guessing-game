# Number-guessing-game
Absolutely! Here’s a **compact, single-space README** you can copy directly:

---

# 🎯 Number Guessing Game in Go

A simple, interactive command-line game where the player guesses a randomly generated number. Features difficulty levels, hints, attempt tracking, and the ability to quit anytime.

## Features

* Random number between 1 and 100
* Difficulty levels: Easy – 10 attempts, Medium – 5 attempts, Hard – 3 attempts
* Hints: Even/odd, divisible by 3 or 5
* Tracks attempts and highest score
* Quit anytime by typing `quit` or `exit`
* Input validation ensures only numbers are accepted

## How to Run

* **Linux/macOS**: `./guessgame`
* **Windows**: `guessgame.exe`

## How to Play

1. Select a difficulty level: 1. Easy (10 attempts), 2. Medium (5 attempts), 3. Hard (3 attempts)
2. Enter your guesses. After each guess, you’ll see if it’s too high/low, hints, and remaining attempts
3. Type `quit` or `exit` anytime to leave the game
4. After a correct guess, your score and the highest score are displayed

## Example Session

```
🎯 Welcome to the Number Guessing Game! 🎯
Select difficulty level:
1. Easy (10 attempts)
2. Medium (5 attempts)
3. Hard (3 attempts)

Enter your choice: 1
Enter your guess: 50
Incorrect! The number is greater than 50
You have 9 attempts left
Hint: The number is even
Enter your guess: quit
You chose to exit the game. Goodbye!
```


