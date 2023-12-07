package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	doorStatus string
	handStatus string

	doorStatusMux sync.Mutex
	handStatusMux sync.Mutex
)

func hand() {
	handStatusMux.Lock()
	handStatus = "in"

	time.Sleep(time.Millisecond * 200)

	handStatus = "out"
	handStatusMux.Unlock()

	wg.Done()
}

func door() {

	doorStatusMux.Lock()
	doorStatus = "close"
	doorStatusMux.Unlock()

	time.Sleep(time.Millisecond * 200)

	handStatusMux.Lock()
	if handStatus == "in" {
		fmt.Println("夾到手了啦！")
	} else {
		fmt.Println("沒夾到喔！")
	}
	handStatusMux.Unlock()

	doorStatusMux.Lock()
	doorStatus = "open"
	doorStatusMux.Unlock()
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go door()
		go hand()
		wg.Wait()
		time.Sleep(time.Millisecond * 200)
	}
}
