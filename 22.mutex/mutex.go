package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		// will work wthout go, but not with go(concurrently)
		// That's why need waitgroup
		wg.Add(1)
		go myPost.increament(&wg)
	}
	wg.Wait() // wait for eg to be done first
	fmt.Println(myPost.views)
}

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) increament(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock() // need to lock before modification only
	p.views += 1
}
