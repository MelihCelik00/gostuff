package main

import (
	"fmt"

	"rsc.io/quote/v4"

	"example.com/greetings"

	"log"
)

func main() {
	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// log.SetFlags(3)

	message, status := greetings.Hello("Melih")

	if status != nil { // nil means no error, success!
		log.Fatal(status)
	}

	fmt.Println(message)
}
