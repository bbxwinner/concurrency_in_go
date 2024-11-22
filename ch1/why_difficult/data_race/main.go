package main

import "fmt"

func concurrencyTest() string {
	var data int
	go func() {
		data++
	}()

	if data == 0 {
		return fmt.Sprintf("The value is %v", data)
	}
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
