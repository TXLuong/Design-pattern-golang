package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go say("world", &wg)
	go say("hello", &wg)
	go say("Mark", &wg)

	wg.Wait()
}
