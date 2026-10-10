# Log Analyzer

A command-line log analysis tool built in Go that reads log files, validates entries, summarizes log levels, and extracts error messages.

## Features

- Read log files line by line.
- Parse entries into structured log records.
- Validate timestamps and log levels (`INFO`, `WARNING`, and `ERROR`).
- Count occurrences of each log level.
- Identify the most frequent log level.
- Extract and number error messages.
- Count malformed log lines while continuing to process valid entries.
- Report file-reading errors.
- Run unit tests for log parsing, file reading, and log-level frequency analysis.

## Log Format

The analyzer expects entries in this format:

```text id="y4f8sm"
YYYY-MM-DD HH:MM:SS LEVEL Message
```

Example:

```text id="4o8uqb"
2026-10-08 08:01:12 INFO Server started on port 8080
2026-10-08 08:11:02 ERROR Database connection failed
2026-10-08 08:15:26 WARNING Slow response detected
```

## Prerequisites

- Go installed on your system.
- A log file containing entries in the expected format.

## How to Run

Navigate to this directory and run:

```bash id="t03wmn"
go run . server.log
```

Replace `server.log` with the path to your own log file.

## Running Tests

Run the test suite with:

```bash id="udf9z2"
go test -v
```

## Concepts Practiced

- File handling with `os.Open()`
- Reading files with `bufio.Scanner`
- Structs and slices
- String parsing with `strings.Fields()` and `strings.Join()`
- Timestamp validation with Go's `time` package
- Maps for counting log levels
- Error handling and multiple return values
- Command-line arguments through `os.Args`
- Unit testing with Go's `testing` package

## Purpose

This project brought together several Go concepts to build a practical file-processing CLI tool while introducing input validation, error handling, and unit testing.
