package main

import (
	"fmt"
	"time"
)

func main() {
	/// Buffer chanel
	emailChan := make(chan string, 100) // 100 will not block, after 100, rest will block
	done := make(chan bool)

	go emailSender(emailChan, done)

	// This is not blocking
	for i := 0; i < 10; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("Done sending")
	close(emailChan) // To close chanel. If not, will deadlock
	<-done
	// emailChan <- "test@gmail.com"
	// fmt.Println(<-emailChan)

	/// Un-buffered chanel
	ch := make(chan int)
	boolCh := make(chan bool)

	go sum(ch, 12, 4)

	result := <-ch //blocking
	fmt.Println(result)

	go boolChanel(boolCh)
	<-boolCh // recieving will block ending main()
}

// <-chan : receive-only
// chan<- : send-only
func emailSender(emailChan <-chan string, done chan<- bool) {
	// emailChan <- "sending" // can't send in recieve-only chanel
	// isDone:= <-done // can't recieve in send-only chanel
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func sum(chanel chan int, num1 int, num2 int) {
	sum := num1 + num2
	chanel <- sum
}

func boolChanel(done chan bool) {
	defer func() { done <- true }() // will end chanel after block is complete
	fmt.Println("Processing....")
}
