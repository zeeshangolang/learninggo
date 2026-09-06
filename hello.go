package main

import (
	"context"
	"fmt"
)

func CountTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()

	return ch
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := CountTo(ctx, 9)

	for v := range ch {
		if v > 5 {
			break
		}

		fmt.Print(v)
	}
}
