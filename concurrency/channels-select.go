package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func() {
		for {
			c1 <- "Every 500 milliseconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}		
	}()

	for {
		msg1 := <- c1 
		fmt.Println(msg1)

		msg2 := <- c2 // this blocks and hence msg1 will not be printed until msg2 is available (after 2 seconds)
		fmt.Println(msg2)
	}
}