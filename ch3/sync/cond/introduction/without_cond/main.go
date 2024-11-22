package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var isTrue atomic.Value
	isTrue.Store(false)

	conditionTrue := func() bool {
		return isTrue.Load().(bool)
	}

	setTrue := func() {
		fmt.Println("isTrue after 3s")
		time.Sleep(3 * time.Second)
		isTrue.Store(true)
	}

	go setTrue()

	for !conditionTrue() {
	}

	fmt.Println("Main End")
}
