package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2) // buffer as 2. It will not block until channel is full
	c <- "Hello"
	c <- "World"
	// c <- "Third message" //this will block and fail becuase channel is full with 2 messages

	msg := <- c
	fmt.Println(msg)

	msg = <- c
	fmt.Println(msg)
}