package main

import (
	"fmt"
	"sync"
)

var (
	counter int64
	mutex   sync.Mutex
)

func main() {
	for i := 0; i < 100; i++ {
		go func() { // Hàm vô danh
			for i := 0; i < 10000; i++ {
				// counter++
				// atomic.AddInt64(&counter, 1)

				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	var c string
	fmt.Scanln(&c)
	fmt.Println(counter)
}
