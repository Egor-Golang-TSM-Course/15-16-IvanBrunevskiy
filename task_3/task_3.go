package task_3

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Notifications interface {
	SendNotification(ctx context.Context, userID int, message string)
}

type Email struct{}

func (e *Email) SendNotification(ctx context.Context, userID int, message string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Email: Context canceled. Stopping sending for user %d\n", userID)
			return
		default:
			fmt.Printf("Message For %d: %s\n", userID, message)
			time.Sleep(time.Second)
		}
	}
}

type SMS struct{}

func (s *SMS) SendNotification(ctx context.Context, userID int, message string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("SMS: Context canceled. Stopping sending for user %d\n", userID)
			return
		default:
			fmt.Printf("Message For %d: %s\n", userID, message)
			time.Sleep(time.Second)
		}
	}
}

func StartTask3() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	emailSender := &Email{}
	smsSender := &SMS{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		emailSender.SendNotification(ctx, 1, "Email!")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		smsSender.SendNotification(ctx, 2, "SMS!")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Cancel context.")
	cancel()

	wg.Wait()
}
