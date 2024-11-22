package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			// closure loop issue was fixed in go1.21
			fmt.Println(salutation) // <1>
		}()
	}
	wg.Wait()
}
