package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {

	log.SetPrefix("grettings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	msg, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
