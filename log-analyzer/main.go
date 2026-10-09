package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type LogEntry struct {
	Timestamp string
	LogLevel  string
	Message   string
}

func readLogFile(filename string) ([]LogEntry,int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	logEntries := []LogEntry{}
	scanner := bufio.NewScanner(file)
	malformedCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		logEntry, err := parseLogLine(line)
		if err != nil {
			malformedCount++
			continue
		}
		logEntries = append(logEntries, logEntry)
	}

	if err := scanner.Err(); err != nil {
		return logEntries, malformedCount, err
	}

	return logEntries, malformedCount, nil
}

func parseLogLine(line string) (LogEntry, error) {
	parts := strings.Fields(line)
	if len(parts) < 4 {
		return LogEntry{}, fmt.Errorf("invalid log line: %s", line)
	}

	switch parts[2] {
	case "INFO", "WARNING", "ERROR":
		// Valid log levels
	default:
		return LogEntry{}, fmt.Errorf("invalid log level: %s", parts[2])
	}

	date := parts[0]
	times := parts[1]
	logLevel := parts[2]
	timestamp := date + " " + times
	message := strings.Join(parts[3:], " ")

	_, err := time.Parse("2006-01-02 15:04:05", timestamp)
	if err != nil {
		return LogEntry{}, fmt.Errorf("invalid timestamp format: %s", timestamp)
	}

	return LogEntry{
		Timestamp: timestamp,
		LogLevel:  logLevel,
		Message:   message,
	}, nil
}

func findMostFrequentLevel(counts map[string]int) (string, int) {
	maxCount := 0
	var mostFrequentLevel string
	for level, count := range counts {
		if count > maxCount {
			maxCount = count
			mostFrequentLevel = level
		}
	}
	return mostFrequentLevel, maxCount
}

func extractErrorMessages(entries []LogEntry) []string {
	errorMessages := []string{}
	for _, entry := range entries {
		if entry.LogLevel == "ERROR" {
			errorMessages = append(errorMessages, entry.Message)
		}
	}
	return errorMessages
}

func printReport(entries []LogEntry, malformedCount int) {

	logLevelCounts := make(map[string]int)
	for _, entry := range entries {
		logLevelCounts[entry.LogLevel]++
	}

	fmt.Println("Log Level Counts:")
	for level, count := range logLevelCounts {
		fmt.Printf("%s: %d\n", level, count)
	}

	mostFrequentLevel, maxCount := findMostFrequentLevel(logLevelCounts)
	fmt.Printf("Most Frequent Log Level: %s (%d occurrences)\n", mostFrequentLevel, maxCount)

	fmt.Printf("Total Valid Entries: %d\n", len(entries))
	fmt.Printf("Malformed Lines: %d\n", malformedCount)

	errorMessages := extractErrorMessages(entries)
	fmt.Println("\nError Messages:")
	for i, msg := range errorMessages {
		fmt.Printf("%d: %s\n", i+1, msg)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <logfile>")
		return
	}

	logFile := os.Args[1]
	logEntries, malformedCount, err := readLogFile(logFile)
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	printReport(logEntries,malformedCount)

}
