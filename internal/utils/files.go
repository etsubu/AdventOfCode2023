package utils

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
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

func LoadLinesFromUrl(url string) ([]string, error) {
	var lines []string
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return lines, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return lines, err
	}
	return strings.Split(string(body), "\n"), nil
}
