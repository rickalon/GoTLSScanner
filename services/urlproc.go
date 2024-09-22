package services

import (
	"context"
	"fmt"
	"time"
)

func OrDone(ctx context.Context, ch <-chan string) <-chan string {
	stream := make(chan string)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Salimos")
				return
			case val, ok := <-ch:
				if !ok {
					return
				}
				select {
				case stream <- val:
				case <-ctx.Done():
					return
				}
			default:
				fmt.Println("Test...")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	return stream
}
