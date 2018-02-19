package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(done chan struct{}, wg * sync.WaitGroup) {
	defer wg.Done()

	select {
	case <- done:
		return
	default:
		fmt.Println("start worker")
		time.Sleep(2 * time.Second)
		fmt.Println("end worker")
	}
}

func main() {
	fmt.Println("start main")

	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go worker(done, wg)

	time.Sleep(1 * time.Second)
	close(done)
	wg.Wait()

	fmt.Println("end main")
}
