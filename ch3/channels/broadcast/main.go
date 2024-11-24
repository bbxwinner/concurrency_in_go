package main

import (
	"fmt"
	"sync"
)

func main() {
	syncChan := make(chan struct{})
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-syncChan
			fmt.Println(id, "beginning")
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	close(syncChan)
	wg.Wait()
}
