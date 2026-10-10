package main

import (
	"testing"
)

func TestReadLogFile(t *testing.T) {
	filename := "server.log"
	entries, malformedCount, err := readLogFile(filename)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}
	if len(entries) == 0 {
		t.Errorf("Expected log entries, but got none")
	}
	if malformedCount == 0 {
		t.Errorf("Expected malformed lines, but got none")
	}
}

func TestParseLogLine(t *testing.T) {
	validLine := "2026-10-08 08:23:45 INFO Server started"
	entry, err := parseLogLine(validLine)
	if err != nil {
		t.Fatalf("Failed to parse valid log line: %v", err)
	}
	if entry.Timestamp != "2026-10-08 08:23:45" || entry.LogLevel != "INFO" || entry.Message != "Server started" {
		t.Errorf("Parsed log entry does not match expected values")
	}
	
	invalidLine := "2026-10-08 08:23:45 INVALID Server started"
	_, err = parseLogLine(invalidLine)	
	if err == nil {
		t.Errorf("Expected error for invalid log level, but got none")
	}

	invalidLine2 := "2026-10-08 08:23:45 INFO"
	_, err = parseLogLine(invalidLine2)
	if err == nil {
		t.Errorf("Expected error for missing message, but got none")
	}

	invalidLine3 := "not-a-date not-a-time INFO Server started"
	_, err = parseLogLine(invalidLine3)
	if err == nil {
		t.Errorf("Expected error for invalid timestamp format, but got none")
	}	

}

func TestMostFrequentLevel(t *testing.T) {

	logLevelCounts := make(map[string]int)
	logLevelCounts["INFO"] = 5
	logLevelCounts["WARNING"] = 3
	logLevelCounts["ERROR"] = 7

	mostFrequentLevel, maxCount := findMostFrequentLevel(logLevelCounts)
	if mostFrequentLevel != "ERROR" || maxCount != 7 {
		t.Errorf("Expected most frequent level to be ERROR with count 7, but got %s with count %d", mostFrequentLevel, maxCount)
	}

	emptyCounts := make(map[string]int)
	mostFrequentLevel, maxCount = findMostFrequentLevel(emptyCounts)
	if mostFrequentLevel != "" || maxCount != 0 {
		t.Errorf("Expected most frequent level to be empty with count 0, but got %s with count %d", mostFrequentLevel, maxCount)
	}

}