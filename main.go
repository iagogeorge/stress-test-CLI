package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Define CLI flags
	url := flag.String("url", "", "URL of the service to test")
	requests := flag.Int("requests", 100, "Total number of requests to perform")
	concurrency := flag.Int("concurrency", 10, "Number of concurrent requests")
	flag.Parse()

	if *url == "" {
		log.Fatal("URL must be provided")
	}
	if *requests <= 0 || *concurrency <= 0 {
		log.Fatal("Requests and concurrency must be positive integers")
	}

	startTime := time.Now()
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, *concurrency)
	var mu sync.Mutex
	statusCounts := make(map[int]int)

	for i := 0; i < *requests; i++ {
		wg.Add(1)
		semaphore <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()
			resp, err := http.Get(*url)
			if err != nil {
				log.Printf("Error making request: %v", err)
				return
			}
			defer resp.Body.Close()
			
			mu.Lock()
			statusCounts[resp.StatusCode]++
			mu.Unlock()
		}()
	}

	wg.Wait()
	elapsedTime := time.Since(startTime)

	fmt.Printf("Test completed in %s\n", elapsedTime)
	fmt.Printf("Total requests: %d\n", *requests)
	fmt.Printf("HTTP 200 OK: %d\n", statusCounts[http.StatusOK])
	fmt.Printf("Other HTTP statuses:\n")
	for code, count := range statusCounts {
		if code != http.StatusOK {
			fmt.Printf("  %d: %d\n", code, count)
		}
	}
}
