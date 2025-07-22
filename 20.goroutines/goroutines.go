package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	/// Regular, will do one by one
	// for i := 0; i < 10; i++ {
	// 	tasks(i)
	// }

	/// Will do parallally, concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1) // wg will not work if its zero or negative
		go tasks(i, &wg)
	}

	wg.Wait() // Wait blocks until the [WaitGroup] counter is zero
}

func tasks(id int, w *sync.WaitGroup) {
	defer w.Done() 
	// defer execute code after completing the block
	// Done decrements the [WaitGroup] counter by one
	fmt.Println("doing task", id)
}
