package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message1, err := greetings.Hello("Andi Sholihin")

	fmt.Println(message1)

	message2, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message2)
}
