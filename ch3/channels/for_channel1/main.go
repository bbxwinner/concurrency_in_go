package main

import "fmt"

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 5; i++ {
			intStream <- i
		}
	}()
	for v := range intStream {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
