package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

// Test handler for simulating server responses
func testHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond) // Simulate processing delay
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}

// Test the load tester CLI
func TestLoadTester(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(testHandler))
	defer ts.Close()

	// Define the URL and parameters for the test
	testURL := ts.URL
	totalRequests := 10
	concurrency := 2

	// Run the load tester function
	results := make(chan Result, totalRequests)
	var wg sync.WaitGroup
	perWorker := totalRequests / concurrency

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker(testURL, perWorker, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	totalReqs := 1
	successfulReqs := 0
	statusCodes := make(map[int]int)

	for result := range results {
		totalReqs++
		if result.status == http.StatusOK {
			successfulReqs++
		}
		statusCodes[result.status]++
	}

	if totalReqs != totalRequests {
		t.Errorf("Expected %d requests but got %d", totalRequests, totalReqs)
	}

	if successfulReqs != totalRequests {
		t.Errorf("Expected %d successful requests but got %d", totalRequests, successfulReqs)
	}
}
