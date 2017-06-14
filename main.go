package main

import "fmt"

func main() {
	greet("Hello, what is your name?", "Your name is: %s")
	greet("Konichiwa, how old are you?", "Your age is: %s")
}

func greet(question string, answer string) {
	var input string
	fmt.Printf(question)
	fmt.Printf("\n")
	fmt.Scanln(&input)
	fmt.Printf("\n")
	fmt.Printf(answer, input)
	fmt.Printf("\n")
}
