// Example data races in Golang concurrency
// run with flag -race to test it
package main


import (
	"fmt"
	"time"
	"sync"
	"runtime"
)

func main() {
	// Sets the maximum number of CPUs that can be executing simultaneously
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

	time.Sleep(time.Second)

	mutex.Lock()
	fmt.Println(counter)
	mutex.Unlock()
}
