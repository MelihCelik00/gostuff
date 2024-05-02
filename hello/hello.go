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
	fmt.Println("-----------------")

	names := []string{"Melih", "Safa", "Çelik"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	for _, value := range messages {
		fmt.Println(value)
	}
}
