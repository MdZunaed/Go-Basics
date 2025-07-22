package main

import "fmt"

func main() {
	messageChanel := make(chan string)

	messageChanel <- "ping"

	msg := <-messageChanel

	fmt.Println(msg)
}