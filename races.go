// Counter never be 1000 for NumCPU > 1
package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var mutex sync.Mutex

	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			counter += 1
			mutex.Unlock()
		}()
	}
	
	fmt.Println(counter)
}
