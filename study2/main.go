package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type data struct {
	err   error
	value int
}

func main() {
	// ここは*dataよりdataの方がコストが小さい
	valueChan := make(chan data)
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
		fmt.Println(v.err)
		fmt.Println(v.value)
		if v.err != nil {
			return
		}
	}
}

func process(i int, valueChan chan<- data) {
	if i == 50 {
		valueChan <- data{err: errors.New("error"), value: 0}
		return
	}
	time.Sleep(1 * time.Second)
	valueChan <- data{err: nil, value: i}
}
