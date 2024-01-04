package task_2

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func process(ctx context.Context, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context is canceled.")
			return
		case data, ok := <-in:
			if !ok {
				return
			}
			result := data * 2
			out <- result
		}
	}
}

func StartTask2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	var wg sync.WaitGroup

	go func() {
		defer close(ch1)
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
	}()

	wg.Add(1)
	go process(ctx, ch1, ch2, &wg)

	wg.Add(1)
	go process(ctx, ch2, ch3, &wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range ch3 {
			fmt.Println("Result:", result)
		}
	}()

	time.Sleep(1 * time.Second)
	if time.Now().Unix()%2 == 0 {
		fmt.Println("Simulating error. Cancelling context.")
		cancel()
	}

	<-ctx.Done()
}
