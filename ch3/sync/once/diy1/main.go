package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Once1 struct {
	done atomic.Uint32
	m    sync.Mutex
}

// 双重检测锁模式
func (o *Once1) Do(f func()) {
	if o.done.Load() == 0 { // 第一次检查（无锁检查），如果之前 f() 已经被完成调用，那么不需要再重复获取锁
		o.m.Lock()              // 加锁
		defer o.m.Unlock()      // f 可能 panic，使用 defer 确保在函数结束时解锁
		if o.done.Load() == 0 { // 第二次检查（加锁后检查），因为在锁阻塞期间，f() 可能已经被完成调用，此时需要再次检查，防止 f 重复执行
			defer o.done.Store(1) // f 可能 panic，使用 defer 确保操作完成后设置 done 的值
			f()
		}
	}
}

func main() {
	once := Once1{}
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
