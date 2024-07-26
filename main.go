package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	status int
}

func worker(url string, requests int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < requests; i++ {
		resp, err := http.Get(url)
		if err != nil {
			results <- Result{status: 0}
			continue
		}
		results <- Result{status: resp.StatusCode}
		resp.Body.Close()
	}
}

func main() {
	url := flag.String("url", "", "URL to load test")
	requests := flag.Int("requests", 1, "Number of total requests")
	concurrency := flag.Int("concurrency", 1, "Number of concurrent requests")
	flag.Parse()

	if *url == "" {
		fmt.Println("URL is required")
		flag.Usage()
		return
	}

	fmt.Println("Load test process started...")

	start := time.Now()
	var wg sync.WaitGroup
	results := make(chan Result, *requests)
	perWorker := *requests / *concurrency

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go worker(*url, perWorker, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalRequests := 0
	successfulRequests := 0
	statusCodes := make(map[int]int)

	for result := range results {
		totalRequests++
		if result.status == 200 {
			successfulRequests++
		}
		statusCodes[result.status]++
	}

	duration := time.Since(start)

	fmt.Printf("Total time taken: %s\n", duration)
	fmt.Printf("Total requests: %d\n", totalRequests)
	fmt.Printf("Successful requests (status 200): %d\n", successfulRequests)
	fmt.Println("Status code distribution:")
	for code, count := range statusCodes {
		fmt.Printf("  %d: %d\n", code, count)
	}
}
