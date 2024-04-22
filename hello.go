package main

import (
	"example/greetings"
	"fmt"
	"log"

	"rsc.io/quote"
)

func safeGreeting(name string) {
	message, err := greetings.Hello(name)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("[greetings] => ")
	log.SetFlags(0)

	fmt.Println("Hello, world !")
	fmt.Println(quote.Go())
	safeGreeting("Clonkk")
	safeGreeting("")
}
