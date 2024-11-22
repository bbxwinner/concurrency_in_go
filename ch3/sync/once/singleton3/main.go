package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Foo struct {
	id   int
	name string
}

var Singleton *Foo
var m sync.Mutex
var flag atomic.Uint32

func GetFooInstance() *Foo {
	if flag.Load() == 0 {
		m.Lock()
		defer m.Unlock()
		if flag.Load() == 0 {
			defer flag.Store(1)
			Singleton = new(Foo)
		}
	}
	return Singleton
}

func main() {
	wg := sync.WaitGroup{}
	addressSet := sync.Map{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fooInstance := GetFooInstance()
			addressSet.Store(fmt.Sprintf("%p", fooInstance), true)
		}()
	}
	wg.Wait()
	addressSet.Range(func(key, value interface{}) bool {
		fmt.Println(key)
		return true
	})
}
