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

// main is the entry point for the executable program.
func main() {
	// Initialize a new buffered reader to efficiently read input from the standard input stream (the console).
	reader := bufio.NewReader(os.Stdin)

	// Define the default range for the guessing game.
	const defaultMin = 1
	const defaultMax = 100
	min := defaultMin
	max := defaultMax

	// Display welcome message and prompt the user for a custom guessing interval.
	fmt.Println("Welcome to the number guessing game!")
	fmt.Printf("You can enter your own interval (e.g., '1-50') or press Enter for the default interval (%d-%d).\n", defaultMin, defaultMax)
	fmt.Print("Your interval choice: ")

	// Read the user's input for the interval choice until a newline character ('\n') is encountered.
	intervalInput, _ := reader.ReadString('\n')
	// Remove leading and trailing whitespace, including the newline character.
	intervalInput = strings.TrimSpace(intervalInput)

	// Process custom interval input if the user entered anything.
	if intervalInput != "" {
		// Split the input string based on the '-' delimiter.
		parts := strings.Split(intervalInput, "-")
		// Check if the input was split into exactly two parts (min and max).
		if len(parts) == 2 {
			// Attempt to convert both parts (after trimming whitespace) into integers.
			minInput, errMin := strconv.Atoi(strings.TrimSpace(parts[0]))
			maxInput, errMax := strconv.Atoi(strings.TrimSpace(parts[1]))

			// Validate the conversion and ensure the minimum is strictly less than the maximum.
			if errMin == nil && errMax == nil && minInput < maxInput {
				// Success: Update the game's range.
				min = minInput
				max = maxInput
			} else {
				// Failure: Invalid number format or min >= max. Inform user and keep the default.
				fmt.Printf("Invalid interval format or range (min >= max). Using the default interval %d-%d.\n", defaultMin, defaultMax)
			}
		} else {
			// Failure: Input was not in the expected 'X-Y' format. Inform user and keep the default.
			fmt.Printf("Invalid interval format. Using the default interval %d-%d.\n", defaultMin, defaultMax)
		}
	}

	// Seed the pseudo-random number generator with the current nanosecond timestamp to ensure different results on each run.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate the secret number: r.Intn(n) generates a number in [0, n-1].
	// To get a number in [min, max], we use r.Intn(max-min+1) + min.
	secretNumber := r.Intn(max-min+1) + min
	// Initialize the counter for the number of guessing attempts.
	attempts := 0

	// Inform the user about the final range chosen for the game.
	fmt.Printf("I'm thinking of a number between %d and %d. Try to guess it.\n", min, max)

	// Start the main game loop. It continues indefinitely until explicitly broken out of (the correct guess is made).
	for {
		fmt.Print("Your guess: ")
		// Read the user's guess input.
		input, _ := reader.ReadString('\n')
		// Clean up the input string.
		input = strings.TrimSpace(input)

		// Attempt to convert the input string to an integer.
		guess, err := strconv.Atoi(input)
		// Check for conversion error (e.g., the user didn't enter a number).
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue // Skip the rest of the loop body and start the next iteration.
		}

		// A valid guess was made, so increment the attempt counter.
		attempts++

		// Compare the guess with the secret number and provide feedback.
		if guess < secretNumber {
			fmt.Printf("Too low! Try a number between %d and %d.\n", min, max)
		} else if guess > secretNumber {
			fmt.Printf("Too high! Try a number between %d and %d.\n", min, max)
		} else {
			// The guess is correct. Display the success message and statistics.
			fmt.Println("Correct! You guessed the number", secretNumber, "in", attempts, "attempts!")
			break // Exit the 'for' loop, ending the game.
		}
	}
}
