package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Once2 struct {
	done atomic.Uint32
	m    sync.Mutex
}

// 双重检测锁模式
func (o *Once2) Do(f func()) {
	if o.done.CompareAndSwap(0, 1) {
		f()
	}
}

func main() {
	once := Once2{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			once.Do(func() {
				time.Sleep(3 * time.Second)
				fmt.Println("executed by goroutine", i)
			})
			fmt.Println("goroutine", i, "finished")
		}(i)
	}
	wg.Wait()
}
