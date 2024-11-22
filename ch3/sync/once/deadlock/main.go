package main

import (
	"sync"
)

func main() {
	var onceA sync.Once
	var initA func()
	initA = func() {
		onceA.Do(initA) // 2. 但是 initA 中再次调用了 onceA.Do，这里会造成死锁
	}
	onceA.Do(initA) // 1. initA 执行完成，onceA.Do 才会往下走
}
