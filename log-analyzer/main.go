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

func main() {

	file, err := os.Open("server.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()
	logEntries := []LogEntry{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		date := parts[0]
		time := parts[1]
		logLevel := parts[2]
		timestamp := date + " " + time
		message := strings.Join(parts[3:], " ")

		logEntry := LogEntry{
			Timestamp: timestamp,
			LogLevel:  logLevel,
			Message:   message,
		}

		logEntries = append(logEntries, logEntry)
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	
	logLevelCounts := make(map[string]int)

	for _, entry := range logEntries {
		logLevelCounts[entry.LogLevel]++
		fmt.Printf("[%s] %s: %s\n", entry.Timestamp, entry.LogLevel, entry.Message)
	}
	

	fmt.Println("\nLog Level Counts:")
	for level, count := range logLevelCounts {

		fmt.Printf("%s: %d\n", level, count)
	}
	
	fmt.Printf("Total Log Entries: %d\n", len(logEntries))


}
