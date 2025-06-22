package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	message := greet(name)
	fmt.Println(message)
}

func greet(name string) string {
	return "Hello, " + name + "! Welcome to Go."
}
