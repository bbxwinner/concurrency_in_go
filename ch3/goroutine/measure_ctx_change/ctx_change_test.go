package main

import (
	"sync"
	"testing"
)

// go test -bench=. -cpu=1 concurrency_in_go/ch3/sync/goroutine/ctx_change/ctx_change_test.go
// goos: linux
// goarch: amd64
// cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
// BenchmarkContextSwitch   6501746               173.1 ns/op
// PASS
// ok      command-line-arguments  1.315s

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin // <1>
		for i := 0; i < b.N; i++ {
			c <- token // <2>
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin // <1>
		for i := 0; i < b.N; i++ {
			<-c // <3>
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer() // <4>
	close(begin)   // <5>
	wg.Wait()
}
