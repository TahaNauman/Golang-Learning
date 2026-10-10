# Go Learning — Mini Projects

A collection of hands-on projects built while learning **Go (Golang)**, progressing from fundamental programming concepts to practical CLI tools and, eventually, backend development.

The goal of this repository is to learn by building, understand Go's standard library, and develop good coding practices through progressively more challenging projects.

## Projects

| # | Project | Description | Concepts Practiced |
|---|---|---|---|
| 1 | [Number Guessing Game](./number-guessing/) | A command-line guessing game with difficulty levels and limited attempts. | Variables, loops, conditionals, functions, random numbers, input handling |
| 2 | [To-Do List](./to-do-list/) | A CLI task manager for adding, viewing, editing, completing, and deleting tasks. | Structs, slices, functions, input validation, CRUD operations |
| 3 | [Expense Tracker](./expense-tracker/) | A command-line application for recording and managing expenses. | Structs, slices, maps, numeric input, aggregation |
| 4 | [Student Grade Manager](./grades-manager/) | A program for managing student grades and calculating statistics. | Functions, slices, averages, minimum and maximum values, multiple return values |
| 5 | [Log Analyzer](./log-analyzer/) | A CLI tool that parses log files, counts log levels, extracts error messages, and tracks malformed entries. | File I/O, structs, error handling, string parsing, maps, timestamps, unit testing |

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) installed on your system.
- A terminal or command prompt.

Verify your installation:

```bash
go version
```

### Running a Project

Navigate to the directory of the project you want to run. For example:

```bash
cd number-guessing
go run .
```

For the Log Analyzer:

```bash
cd log-analyzer
go run . server.log
```

To run the Log Analyzer's tests:

```bash
go test -v
```

Check each project's README for its specific instructions and requirements.

## Learning Roadmap

This repository is an ongoing learning journey. Future projects will explore topics such as:

- Concurrent programming with goroutines and channels
- HTTP servers and REST APIs
- JSON handling and database integration
- Automated testing and software quality
- Docker, CI/CD, and cloud deployment

These topics will be added as they are explored and implemented.

