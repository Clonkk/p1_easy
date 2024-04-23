package main

import (
	"example/greetings"
	"fmt"
	"log"

	"rsc.io/quote"
)

func safeGreeting(names []string) {
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	// to the console.
	for _, message := range messages {
		fmt.Println(message)
	}
}

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("[greetings] => ")
	log.SetFlags(0)

	fmt.Println("Hello, world !")
	fmt.Println(quote.Go())
	names := []string{"Clonkk", "", "Glados"}
	safeGreeting(names)
}
