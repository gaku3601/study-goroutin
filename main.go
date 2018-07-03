package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	workernum := 100
	errChan := make(chan error)
	valueChan := make(chan int)
	// 起動groutin制御(3つまで同時起動)
	limit := make(chan int, 3)
	go func() {
		for i := 0; i < workernum; i++ {
			limit <- 1
			go func(i int) {
				process(i, errChan, valueChan)
				<-limit
			}(i)
		}
	}()
	for i := 0; i < workernum; i++ {
		select {
		case value := <-valueChan:
			// 処理が成功した場合の処理
			fmt.Println(value)
		case err := <-errChan:
			// 処理が失敗した場合の処理
			fmt.Println(err)
			return
		}
	}
}

func process(i int, errChan chan<- error, valueChan chan<- int) {
	// 50がきたら強制的にエラーをはく
	if i == 50 {
		errChan <- errors.New("error")
		return
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("処理中:%d\n", i)
	valueChan <- i
}

/*
func main() {
	errChan := make(chan error, 10)
	valueChan := make(chan int, 10)
	var s, e time.Time
	s = time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			process(j, errChan, valueChan)
		}(i)
	}
	go func() {
		wg.Wait()
		defer close(errChan)
		defer close(valueChan)
	}()

	for v := range errChan {
		fmt.Println(v)
	}
	for v := range valueChan {
		fmt.Println(v)
	}

		LOOP:
			for {
				select {
				case value := <-valueChan:
					// 処理が成功した場合の処理
					fmt.Println(value)
				case err := <-errChan:
					// 処理が失敗した場合の処理
					fmt.Println(err)
				default:
					break LOOP
				}
			}
	e = time.Now()
	fmt.Printf("処理完了 : %v Seconds\n", (e.Sub(s)).Seconds())
}

func process(i int, errChan chan<- error, valueChan chan<- int) {
	if i == 5 {
		errChan <- errors.New("error")
		return
	}
	time.Sleep(3 * time.Second)
	fmt.Printf("処理中:%d\n", i)
	valueChan <- i
}
*/
