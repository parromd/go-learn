package main

import (
	"fmt"
)


func main() {

	messages := make(chan string)

	go func() { messages <- "real msg" }()

	msg := <-messages
	if msg == "ping" {
		fmt.Println("pong")
	} else {
		fmt.Println(msg)
	}
}
