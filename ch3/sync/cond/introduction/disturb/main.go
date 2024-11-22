package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var c = sync.NewCond(&sync.Mutex{})

	var isTrue atomic.Value
	isTrue.Store(false)

	conditionTrue := func() bool {
		return isTrue.Load().(bool)
	}

	setTrue := func() {
		fmt.Println("isTrue after 3s")
		time.Sleep(3 * time.Second)
		isTrue.Store(true)
		c.Signal()
	}

	go setTrue()

	disturb := func() {
		randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
		nseconds := randSource.Float64() * 3
		time.Sleep(time.Duration(nseconds) * time.Second)
		fmt.Println("Disturb at", nseconds, "s")
		c.L.Lock()
		c.Signal()
		c.L.Unlock()
	}
	go disturb()
	go disturb()
	go disturb()
	go disturb()
	go disturb()
	go disturb()
	go disturb()
	go disturb()
	go disturb()

	c.L.Lock()             // 重要，因为调用 Wait 会进行解锁
	for !conditionTrue() { // 使用 for 避免虚假唤醒
		c.Wait() // 解锁并进入暂停状态，main goroutine 会阻塞等待 Signal/Broadcast 调用
	}
	c.L.Unlock() // 重要，因为调用 Wait 会进行加锁，否则下次 c.L.Lock() 的调用会死锁

	fmt.Println("Main End")
}
