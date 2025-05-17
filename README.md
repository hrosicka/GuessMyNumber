# Tutorial: "Guess My Number" Game in Go (For Complete Beginners)

Welcome to this tutorial that will guide you through the basics of the Go programming language by creating a simple number guessing game.
Go is a modern, compiled programming language designed at Google.
It's known for its simplicity, efficiency, and strong support for concurrent programming.

## What will you learn?

During this tutorial, you will become familiar with the following Go language concepts:

* Basic structure of a Go program (`package main`, `import`, `func main`).
* Importing packages (`fmt`, `math/rand`, `os`, `strconv`, `strings`, `time`).
* Working with variables and data types (especially `int` and `string`).
* Input and output using the `fmt` package.
* Reading user input using the `bufio` and `os` packages.
* Converting strings to numbers using the `strconv` package.
* Using conditional statements (`if`, `else if`, `else`).
* Using loops (`for`).
* Generating random numbers using the `math/rand` package.

## 1. Installing Go

For this tutorial, you don't need any prior knowledge of Go.
You should have Go installed on your computer.
If you don't have it yet, you can download and install it from the [official Go website](https://go.dev/dl/).

After installation, open your terminal and check the version:

```bash
go version
```

This should print the installed Go version.

---

## 2. Recommended Editor (VS Code)

I recommend using [Visual Studio Code](https://code.visualstudio.com/) and installing the “Go” extension by the Go Team at Google. This will make writing and debugging Go code much easier.

1.  Open VS Code.
2.  Click on the square icon on the left sidebar (Extensions).
3.  In the search bar, type "Go".
4.  Find the extension named "Go" by "Go Team at Google" and click the "Install" button.

After installing this extension, developing in Go within VS Code will be much more convenient.

---

## 3. Project Structure

Your main file should be named, for example, `Guess.go`.

---

## 4. How to Run the Program

1. Make sure you are in the folder with this file.
2. Open your terminal and run:

```bash
go run Guess.go
```

3. Follow the instructions in the program.

---

## 5. Tips and Explanations

- Every line of code is commented to help you understand what’s happening.
- If you want to try your own extension (for example, a limited number of guesses or hints), try modifying the code after completing this basic version.

