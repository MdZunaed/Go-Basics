package main

import (
	"fmt"
	"time"
)

func main() {

	/// Regular, will do one by one
	// for i := 0; i < 10; i++ {
	// 	tasks(i)
	// }

	/// Will do parallally, concurrently
	for i := 0; i < 10; i++ {
		go tasks(i)
	}

	// To stop being end main function. as there is no more operation
	time.Sleep(time.Second * 2)
}

func tasks(id int) {
	fmt.Println("doing task", id)
}
