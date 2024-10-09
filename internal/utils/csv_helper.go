package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadAccountsFromCSV(filename string) ([][]string, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %v", err)
	}
	return records, nil
}
