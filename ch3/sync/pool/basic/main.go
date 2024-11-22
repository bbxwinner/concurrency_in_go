package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			foo := new(struct {
				name string
				id   int
			})
			fmt.Printf("Creating new instance: %p\n", foo)
			return foo
		},
	}

	instance1 := myPool.Get() // <1>
	fmt.Printf("1st myPool.Get: %p\n", instance1)
	instance2 := myPool.Get() // <1>
	fmt.Printf("2nd myPool.Get: %p\n", instance2)
	myPool.Put(instance2) // <2>
	fmt.Println("Put into myPool again")
	instance2 = myPool.Get() // <3>
	fmt.Printf("3rd myPool.Get: %p\n", instance2)
}
