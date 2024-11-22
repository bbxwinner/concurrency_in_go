package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := make([]interface{}, 0, 10) // <2>
	msgCh := make(chan struct{}, 2)
	m := sync.Mutex{}

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		m.Lock()
		q0 := queue[0]
		queue = queue[1:] // <9>
		m.Unlock()
		fmt.Println("Removed from queue", q0)
		<-msgCh
	}

	for i := 0; i < 10; i++ {
		msgCh <- struct{}{}
		fmt.Println("Adding to queue", i)
		m.Lock()
		queue = append(queue, i)
		m.Unlock()
		go removeFromQueue(1 * time.Second) // <6>
	}
	fmt.Println(queue)
}
