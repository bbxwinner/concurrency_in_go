package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer func() {
			close(intStream)
			fmt.Fprintln(&stdoutBuff, "Producer closed the channel.")
		}()
		for i := 0; i < 100; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %d\n", integer)
	}
}
