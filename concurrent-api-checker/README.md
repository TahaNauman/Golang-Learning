# Concurrent API Checker

A command-line tool built in Go that checks multiple URLs concurrently and reports each one's HTTP status code and response time.

## Overview

This project is a hands-on introduction to concurrency in Go. Instead of checking URLs one at a time, it launches a goroutine per URL and collects the results through a channel, so slow endpoints don't block the others.

## Features

- Check multiple URLs concurrently using goroutines.
- Time each request and report the response duration.
- Print each result as soon as it arrives.
- Count successful (2xx) responses and failed checks.
- Use an HTTP client with a 5-second timeout.
- Built with the standard library only.

## Concepts Learned

- Launching goroutines with the `go` keyword.
- Waiting for goroutines with `sync.WaitGroup`.
- Collecting results through a buffered channel.
- Ranging over a channel and closing it once all workers are done.
- Making HTTP requests with `net/http` and setting a client timeout.
- Measuring elapsed time with `time.Now`, `time.Since`, and `time.Duration`.
- Returning structured results with a `Result` struct.
- Testing HTTP handlers with `httptest`.

## Prerequisites

- Go 1.22 or later installed on your system.
- An internet connection, since the default URLs are live endpoints.

## How to Run

Navigate to this directory and run:

```bash
go run .
```

The URLs to check are defined in `main.go`. Edit that slice to check your own endpoints.

## Running Tests

```bash
go test -v
```

The test suite covers:

- `TestCheckURLSuccess` — a 200 response with no request error.
- `TestCheckURLNotFound` — a 404 response with no request error.
- `TestCheckURLRequestError` — an invalid URL that produces a request error.

## Example Output

```text
Concurrent API Checker
URL: https://httpbin.org/status/200
Response Time: 312.4ms
Status Code: 200

URL: https://example.com
Response Time: 428.7ms
Status Code: 200

URL: https://httpbin.org/status/404
Response Time: 502.1ms
Status Code: 404

----- Summary -----
Successful checks: 2
Failed checks: 1
```

Results arrive in the order the requests finish, so the order may vary between runs.

## A Note on 404 Responses

A 404 is a valid HTTP response, so it produces no request error (the `Err` field stays `nil`). However, it is not a 2xx status, so the checker counts it as an unsuccessful check. Only request failures and non-2xx statuses are counted as failed.
