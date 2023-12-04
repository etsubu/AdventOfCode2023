package utils

import (
	"bufio"
	"os"
)

func LoadLinesFromFile(file string) ([]string, error) {
	var lines []string
	readFile, err := os.Open(file)
	if err != nil {
		return lines, err
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines, nil
}
