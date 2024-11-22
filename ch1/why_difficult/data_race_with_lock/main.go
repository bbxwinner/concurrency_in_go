package main

import (
	"fmt"
	"sync"
)

func concurrencyTest() string {
	var memoryAccess sync.Mutex

	var data int
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 {
		return fmt.Sprintf("The value is %v", data)
	}
	memoryAccess.Unlock()
	return "None"
}

func main() {
	var m = make(map[string]int)
	for i := 0; i < 1000000; i++ {
		m[concurrencyTest()]++
	}
	for k, v := range m {
		fmt.Println(k, "occurs", v)
	}
}
