package main

import (
	"fmt"
	"math"
	"sync"
)

// Pi 结构体
type Pi struct {
	mu    sync.Mutex
	value float64
}

// Value 返回 Pi 的值
func (p *Pi) Value() float64 {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.value
}

// Add 增加 Pi 的值
func (p *Pi) Add(val float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.value += val
}

// FLOOR 函数
func FLOOR(x float64) int64 {
	return int64(math.Floor(x))
}

// CalculatePi calculates digits of Pi between the begin and end
// place.
//
// Internally, CalculatePi will create FLOOR((end-begin)/2) concurrent
// processes which recursively call CalculatePi. Synchronization of
// writes to pi are handled internally by the Pi struct.
func CalculatePi(begin, end int64, pi *Pi) {
	if end-begin <= 1 {
		// 简单的计算 Pi 的一部分
		pi.Add(4.0 * math.Pow(-1, float64(begin)) / (2*float64(begin) + 1))
		return
	}

	mid := begin + (end-begin)/2
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		CalculatePi(begin, mid, pi)
	}()

	go func() {
		defer wg.Done()
		CalculatePi(mid, end, pi)
	}()

	wg.Wait()
}

func main() {
	var pi Pi = Pi{value: 0.0}
	CalculatePi(0, 10000, &pi)
	fmt.Println(pi.Value())
}
