package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Err        error
}

func checkURL(client *http.Client, url string) Result {
	start := time.Now()

	response, err := client.Get(url)
	elapsed := time.Since(start)

	if err != nil {
		return Result{
			URL:      url,
			Duration: elapsed,
			Err:      err,
		}
	}

	defer response.Body.Close()

	return Result{
		URL:        url,
		StatusCode: response.StatusCode,
		Duration:   elapsed,
		Err:        nil,
	}
}

func main() {

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	fmt.Println("Concurrent API Checker")

	url := []string{
		"https://example.com",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
	}

	results := make(chan Result, len(url))

	var wg sync.WaitGroup

	for _, u := range url {

		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			result := checkURL(client, u)
			results <- result
		}(u)
	}
	
	go func() {
		wg.Wait()
		close(results)
	}()

	successfulCount := 0
	failureCount := 0
	for result := range results {
		if result.Err != nil {
			fmt.Printf("Error checking %s: %v\n", result.URL, result.Err)
			failureCount++
			continue
		}

		fmt.Printf("URL: %s\n", result.URL)
		fmt.Printf("Response Time: %v\n", result.Duration)
		fmt.Printf("Status Code: %d\n\n", result.StatusCode)
		if result.StatusCode >= 200 && result.StatusCode < 300 {
			successfulCount++
		} else {
			failureCount++
		}
	}

	fmt.Println("----- Summary -----")
	fmt.Printf("Successful checks: %d\n", successfulCount)
	fmt.Printf("Failed checks: %d\n", failureCount)
}
