package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LogEntry struct {
	Timestamp string
	LogLevel  string
	Message   string
}

func readLogFile(filename string) ([]LogEntry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	logEntries := []LogEntry{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		logEntry, err := parseLogLine(line)
		if err != nil {
			fmt.Println("Error parsing log line:", err)
			continue
		}
		logEntries = append(logEntries, logEntry)
	}

	if err := scanner.Err(); err != nil {
		return logEntries, err
	}

	return logEntries, nil
}

func parseLogLine(line string) (LogEntry, error) {
	parts := strings.Fields(line)
	if len(parts) < 4 {
		return LogEntry{}, fmt.Errorf("invalid log line: %s", line)
	}

	date := parts[0]
	time := parts[1]
	logLevel := parts[2]
	timestamp := date + " " + time
	message := strings.Join(parts[3:], " ")

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

func printReport(entries []LogEntry) {

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

	fmt.Printf("Toal Number of Entries are %d", len(entries))

	errorMessages := extractErrorMessages(entries)
	fmt.Println("\nError Messages:")
	for _, msg := range errorMessages {
		fmt.Println(msg)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <logfile>")
		return
	}

	logFile := os.Args[1]
	logEntries, err := readLogFile(logFile)
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	printReport(logEntries)

}
