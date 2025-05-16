package main

// 'package main' declares that this code is an executable program.
// In Go, every executable program belongs to the 'main' package.

import (
	"bufio"     // Provides functions for buffered I/O, which is an efficient way to read input, e.g., from the user.
	"fmt"       // Implements formatted I/O (input/output), similar to 'printf' and 'scanf' in C. We use it to print text to the console and read input.
	"math/rand" // Contains functions for generating pseudo-random numbers.
	"os"        // Provides operating system-independent functionality, such as interacting with standard input and output.
	"strconv"   // Used for converting between different data types, such as strings to numbers.
	"strings"   // Offers functions for manipulating strings (text).
	"time"      // Provides functions for working with time; here, we use it to seed the random number generator.
)

func main() {
	// 'func main()' is the main function that is executed when you run the program.

	// Setup for reading input from the user via the console.
	reader := bufio.NewReader(os.Stdin)

	// Defining the default guessing interval.
	min := 1
	max := 100

	// Initial messages for the user.
	fmt.Println("Welcome to the number guessing game!")
	fmt.Printf("You can enter your own interval (e.g., '1-50') or press Enter for the default interval (%d-%d).\n", min, max)
	fmt.Print("Your interval choice: ")

	// Reading the user's input for the interval.
	intervalInput, _ := reader.ReadString('\n')
	intervalInput = strings.TrimSpace(intervalInput)

	// Processing input for a custom interval, if provided.
	if intervalInput != "" {
		// Splitting the entered string by the '-' character.
		parts := strings.Split(intervalInput, "-")
		// We expect two parts: the minimum and the maximum.
		if len(parts) == 2 {
			// Attempting to convert both parts to integers.
			minInput, errMin := strconv.Atoi(strings.TrimSpace(parts[0]))
			maxInput, errMax := strconv.Atoi(strings.TrimSpace(parts[1]))
			// If the conversion is successful and the minimum is less than the maximum, we use the provided interval.
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
	// Generate the secret number within the chosen (or default) interval.
	secretNumber := r.Intn(max-min+1) + min
	// Initialize the attempt counter.
	attempts := 0

	// Informing the user about the game range.
	fmt.Printf("I'm thinking of a number between %d and %d. Try to guess it.\n", min, max)

	// The main game loop. It will continue until the player guesses the correct number.
	for {
		fmt.Print("Your guess: ")
		// Reading the user's guess.
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Attempting to convert the guess to an integer.
		guess, err := strconv.Atoi(input)
		// If the conversion fails, inform the user and continue to the next iteration of the loop.
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// Increment the attempt counter.
		attempts++

		// Comparing the guess with the secret number and providing feedback.
		if guess < secretNumber {
			fmt.Printf("Too low! Try a number between %d and %d.\n", min, max)
		} else if guess > secretNumber {
			fmt.Printf("Too high! Try a number between %d and %d.\n", min, max)
		} else {
			// If the player guesses correctly, display a message and break out of the loop.
			fmt.Println("Correct! You guessed the number", secretNumber, "in", attempts, "attempts!")
			break // Exit the 'for' loop.
		}
	}
}
