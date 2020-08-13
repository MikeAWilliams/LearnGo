package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	var wg sync.WaitGroup
	wg.Add(1)

	const messagesToSend = 1000000

	go func() {
		wg.Wait()
		for i := 0; i < messagesToSend; i++ {
			c1 <- struct{}{}
		}
		close(c1)
	}()

	go func() {
		wg.Wait()
		for i := 0; i < messagesToSend; i++ {
			c2 <- struct{}{}
		}
		close(c2)
	}()

	var c1Count, c2Count int
	wg.Done()
	for {
		select {
		case _, ok := <-c1:
			if ok {
				c1Count++
			} else {
				c1 = nil
			}

		case _, ok := <-c2:
			if ok {
				c2Count++
			} else {
				c2 = nil
			}

		}
		if nil == c1 && nil == c2 {
			break
		}
	}

	fmt.Printf("c1Misses: %d\nc2Misses: %d\n", messagesToSend-c1Count, messagesToSend-c2Count)
}
