package main

import (
	"fmt"
	"sync"
)

type Foo struct {
	id   int
	name string
}

var Singleton *Foo
var m sync.Mutex

func GetFooInstance() *Foo {
	m.Lock()
	defer m.Unlock()
	if Singleton == nil {
		Singleton = new(Foo)
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
