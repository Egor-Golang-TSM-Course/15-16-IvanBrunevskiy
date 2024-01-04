package task_1

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func request(ctx context.Context, url string, ch chan<- string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ch <- fmt.Sprintf("Error create request: %s", url)
		return
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("Server error: %s", url)
		return
	}
	defer resp.Body.Close()

	ch <- fmt.Sprintf("URL: %s, Status Code: %d", url, resp.StatusCode)
}

func StartTask1() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resultChanel := make(chan string)

	urls := []string{
		"http://google.com",
		"http://yandex.ru",
		"http://facebook.com",
	}

	for _, url := range urls {
		time.Sleep(1 * time.Second)
		go request(ctx, url, resultChanel)
	}

	for i := 0; i < len(urls); i++ {
		select {
		case result := <-resultChanel:
			fmt.Println(result)
		case <-ctx.Done():
			fmt.Println("Timeout exceeded.")
			return
		}
	}
}
