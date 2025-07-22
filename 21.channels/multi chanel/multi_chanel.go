package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 12
	}()

	go func() {
		chan2 <- "test"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Recieved data from chanel 1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Recieved data from chanel 2", chan2Val)
		}
	}
}
