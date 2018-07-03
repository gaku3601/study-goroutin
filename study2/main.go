package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	valueChan := make(chan int)
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			process(i, valueChan)
		}(i)
	}

	go func() {
		wg.Wait()
		close(valueChan)
	}()
	for v := range valueChan {
		fmt.Println(v)
	}
}

func process(i int, valueChan chan<- int) {
	time.Sleep(1 * time.Second)
	valueChan <- i
}
