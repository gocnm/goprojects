package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	users := []string{"Noor", "Mike", "Peter"}
	messages, error := greetings.Hellos(users)
	
	if(error != nil) {
		log.Fatal(error)
	}

	for person, message := range messages {		
		fmt.Println(person, ":", message)
	}
}
