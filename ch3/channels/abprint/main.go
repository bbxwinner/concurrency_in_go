package main

// print ababababa... using two goroutines

import (
	"fmt"
	"time"
)

// func main() {
// 	f := func(s string) {
// 		fmt.Print(s)
// 		time.Sleep(1 * time.Second)
// 	}

// 	aChan := make(chan struct{})
// 	bChan := make(chan struct{})
// 	go func() {
// 		for {
// 			<-aChan
// 			f("a")
// 			bChan <- struct{}{}
// 		}
// 	}()
// 	go func() {
// 		for {
// 			<-bChan
// 			f("b")
// 			aChan <- struct{}{}
// 		}
// 	}()

// 	aChan <- struct{}{}
// 	time.Sleep(10 * time.Second)
// 	fmt.Println()
// }

func main() {
	f := func(s string) {
		fmt.Print(s)
		time.Sleep(1 * time.Second)
	}

	abChan := make(chan struct{})
	go func() {
		for {
			<-abChan
			f("a")
			abChan <- struct{}{}
		}
	}()
	go func() {
		for {
			<-abChan
			f("b")
			abChan <- struct{}{}
		}
	}()
	abChan <- struct{}{}
	time.Sleep(10 * time.Second)
	fmt.Println()
}
