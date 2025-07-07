package main

import (
	"fmt"
	"time"
)

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}

func doubler(ch <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range ch {
			time.Sleep(500 * time.Millisecond)
			out <- v * 2
		}
		close(out)
	}()

	return out
}

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	reader(doubler(writer()))
}
