package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	min := 1   // Default lower bound of the interval
	max := 100 // Default upper bound of the interval

	fmt.Println("Welcome to the number guessing game!")
	fmt.Printf("You can enter your own interval (e.g., '1-50') or press Enter for the default interval (%d-%d).\n", min, max)
	fmt.Print("Your interval choice: ")
	intervalInput, _ := reader.ReadString('\n')
	intervalInput = strings.TrimSpace(intervalInput)

	// Processing input for a custom interval, if provided.
	if intervalInput != "" {
		parts := strings.Split(intervalInput, "-")
		if len(parts) == 2 {
			minInput, errMin := strconv.Atoi(strings.TrimSpace(parts[0]))
			maxInput, errMax := strconv.Atoi(strings.TrimSpace(parts[1]))
			// Validation of the entered interval.
			if errMin == nil && errMax == nil && minInput < maxInput {
				min = minInput
				max = maxInput
			} else {
				fmt.Printf("Invalid interval format. Using the default interval %d-%d.\n", 1, 100)
			}
		} else {
			fmt.Printf("Invalid interval format. Using the default interval %d-%d.\n", 1, 100)
		}
	}

	// Initialize the random number generator with a unique seed.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate the secret number within the defined interval (inclusive).
	secretNumber := r.Intn(max-min+1) + min
	attempts := 0 // Counter for the player's attempts.

	fmt.Printf("I'm thinking of a number between %d and %d. Try to guess it.\n", min, max)

	// Main game loop.
	for {
		fmt.Print("Your guess: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		// Handling non-numeric input.
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		attempts++ // Increment the attempt counter.

		// Providing feedback to the player.
		if guess < secretNumber {
			fmt.Printf("Too low! Try a number between %d and %d.\n", min, max)
		} else if guess > secretNumber {
			fmt.Printf("Too high! Try a number between %d and %d.\n", min, max)
		} else {
			fmt.Println("Correct! You guessed the number", secretNumber, "in", attempts, "attempts!")
			break // End the game after a correct guess.
		}
	}
}
