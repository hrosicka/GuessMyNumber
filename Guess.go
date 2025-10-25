package main

// Package main declares that this code is an executable program.
// In Go, every executable program must belong to the 'main' package and contain a 'main' function.

import (
	"bufio"     // Provides functions for buffered I/O, typically used for efficient reading of standard input (user input).
	"fmt"       // Implements formatted I/O (input/output), used for printing messages to the console and handling output.
	"math/rand" // Contains functions for generating pseudo-random numbers.
	"os"        // Provides a platform-independent interface to operating system functionality, primarily used here for standard input (os.Stdin).
	"strconv"   // Used for converting between strings and numeric types (e.g., converting user input from a string to an integer).
	"strings"   // Offers functions for manipulating strings, like trimming whitespace or splitting a string.
	"time"      // Provides functionality for measuring and displaying time; essential for seeding the random number generator.
)

// Define the default range for the guessing game as constants.
const defaultMin = 1
const defaultMax = 100

// main is the entry point for the executable program.
func main() {
	// Initialize a new buffered reader to efficiently read input from the standard input stream (the console).
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Number Guessing Game!")

	// The outer game loop: this controls the restart functionality.
	// It continues indefinitely until the user explicitly chooses to quit.
	for {
		// --- GAME SETUP AND INTERVAL SELECTION ---

		// Reset the guessing range to the default at the start of each new game.
		min := defaultMin
		max := defaultMax

		// Prompt the user for a custom guessing interval.
		fmt.Printf("\nYou can enter your own interval (e.g., '1-50') or press Enter for the default interval (%d-%d).\n", defaultMin, defaultMax)
		fmt.Print("Your interval choice: ")

		// Read the user's input for the interval choice.
		intervalInput, _ := reader.ReadString('\n')
		intervalInput = strings.TrimSpace(intervalInput)

		// Process custom interval input if provided.
		if intervalInput != "" {
			parts := strings.Split(intervalInput, "-")
			if len(parts) == 2 {
				// Attempt to convert both parts into integers.
				minInput, errMin := strconv.Atoi(strings.TrimSpace(parts[0]))
				maxInput, errMax := strconv.Atoi(strings.TrimSpace(parts[1]))

				// Validate the conversion and range.
				if errMin == nil && errMax == nil && minInput < maxInput {
					min = minInput
					max = maxInput
				} else {
					fmt.Printf("Invalid interval format or range (min >= max). Using the default interval %d-%d.\n", defaultMin, defaultMax)
				}
			} else {
				fmt.Printf("Invalid interval format. Using the default interval %d-%d.\n", defaultMin, defaultMax)
			}
		}

		// Initialize the random number generator with a unique seed based on the current time.
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// Generate the secret number within the chosen range [min, max].
		secretNumber := r.Intn(max-min+1) + min
		// Initialize the attempt counter for this round.
		attempts := 0

		// Inform the user about the final game range.
		fmt.Printf("\nI'm thinking of a number between %d and %d. Try to guess it.\n", min, max)

		// --- MAIN GUESSING LOOP ---
		// This inner loop handles the actual guessing process for a single game.
		for {
			fmt.Print("Your guess: ")
			// Read and clean the guess input.
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// Attempt to convert the input to an integer.
			guess, err := strconv.Atoi(input)

			// Handle non-numeric input.
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}

			attempts++

			// Provide feedback to the user.
			if guess < secretNumber {
				fmt.Printf("Too low! The number is between %d and %d.\n", min, max)
			} else if guess > secretNumber {
				fmt.Printf("Too high! The number is between %d and %d.\n", min, max)
			} else {
				// Correct guess: End the game round.
				fmt.Println("\n*************************************************")
				fmt.Println("🥳 Correct! You guessed the number", secretNumber, "in", attempts, "attempts!")
				fmt.Println("*************************************************")
				break // Exit the INNER guessing loop.
			}
		}

		// --- GAME RESTART PROMPT ---
		// This section asks the user if they want to play another game.

		fmt.Print("Do you want to play again? (y/n): ")
		// Read the user's choice to continue or quit.
		restartInput, _ := reader.ReadString('\n')
		restartInput = strings.TrimSpace(strings.ToLower(restartInput)) // Convert to lowercase for easier check.

		// Check the user's input.
		if restartInput == "y" || restartInput == "yes" {
			// If 'y' or 'yes', the outer 'for' loop continues to the next iteration (a new game).
			fmt.Println("Starting a new game...")
			continue
		} else {
			// If any other input (like 'n' or anything else), break the outer loop and end the program.
			fmt.Println("Thanks for playing! Goodbye.")
			break // Exit the OUTER game loop, ending func main.
		}
	}
}
