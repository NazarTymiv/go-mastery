package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWork() {
	ch := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
	case <-ctx.Done():
		fmt.Println("Timeout, working over 3 seconds")
	}
}

func main() {
	predictableTimeWork()
}
