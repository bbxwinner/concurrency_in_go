package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
		close(stringStream)
	}()
	s, ok := <-stringStream
	fmt.Println(s, ok)

	s, ok = <-stringStream
	fmt.Println(s, ok)

	s, ok = <-stringStream
	fmt.Println(s, ok)
}
