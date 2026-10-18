package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)



func checkURL(client *http.Client, url string) (*http.Response, error) {
	return client.Get(url)
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

	var wg sync.WaitGroup

	for _, u := range url {

		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			fmt.Println("Checking:", u)
			start := time.Now()
			response, err := checkURL(client, u)
			elapsed := time.Since(start)
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}

			response.Body.Close()

			fmt.Printf("Response Time: %v\n", elapsed)
			fmt.Printf("Status Code: %d\n", response.StatusCode)
		}(u)
	}

	wg.Wait()
}
