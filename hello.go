package main

import (
	"example/greetings"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world !")
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("Test 1"))
}
