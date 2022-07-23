package main

import (
	"fmt"
	"log"

	"tutorial/greetings"

	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

	message, err := greetings.Hello("Littlesquid")
	fmt.Println(message)

	// message1, err1 := greetings.Hello("")
	// if err1 != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message1)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
