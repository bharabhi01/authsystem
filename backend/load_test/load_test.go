package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func main() {
	const (
		totalRequests = 1000
		concurrent    = 1000
		apiURL        = "http://localhost:8080/check"
	)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrent)

	start := time.Now()

	for i := 1; i <= totalRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Limit concurrent requests
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Send request
			data := url.Values{}
			data.Set("username", fmt.Sprintf("user%d", id))

			resp, err := http.Post(apiURL, "application/x-www-form-urlencoded",
				strings.NewReader(data.Encode()))

			if err != nil {
				fmt.Printf("Request %d failed: %v\n", id, err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Request %d: Status %d\n", id, resp.StatusCode)
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("\nCompleted %d requests in %v\n", totalRequests, duration)
	fmt.Printf("Requests per second: %.2f\n", float64(totalRequests)/duration.Seconds())
}
