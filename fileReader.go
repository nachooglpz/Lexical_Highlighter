package main

import (
	"bufio"
	"log"
	"os"
)

// https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
func readLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Dont forget to close the file!!!
	defer file.Close()

	var lines []string
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		lines = append(lines, line)
	}

	return lines, scanner.Err()
}