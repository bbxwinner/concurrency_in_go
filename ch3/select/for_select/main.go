package main

import (
	"fmt"
	"time"
)

func main() {
	doneChan := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(doneChan)
	}()

	workCounter := 0
	done := false
	for {
		select {
		case <-doneChan:
			done = true
		default:
		}
		if done {
			break
		}
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
