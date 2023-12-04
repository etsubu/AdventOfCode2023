package main

import (
	"etsubu.com/aoc2023/internal/utils"
	"log"
	"strconv"
)

func MatchDigit(c uint8) int {
	if c >= '0' && c <= '9' {
		i, _ := strconv.Atoi(string(c))
		return i
	}
	return -1
}

func MatchTextualNumeric(value string, index int) int {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, number := range numbers {
		if len(number)+index > len(value) {
			continue
		}
		if value[index:index+len(number)] == number {
			return i + 1
		}
	}
	return -1
}

func MatchAnyFormatNumeric(value string, index int) int {
	if i := MatchDigit(value[index]); i != -1 {
		return i
	}
	return MatchTextualNumeric(value, index)
}

func part1(lines []string) int {
	lines, err := utils.LoadLinesFromFile("day1.txt")
	if err != nil {
		log.Panic(err.Error())
	}
	sum := 0
	for _, line := range lines {
		i := 0
		j := len(line) - 1
		for ; MatchDigit(line[i]) == -1 && i < len(line); i++ {
		}
		for ; MatchDigit(line[j]) == -1 && j >= 0; j-- {
		}
		value, err := strconv.Atoi(string(line[i]) + string(line[j]))
		if err != nil {
			log.Panic(err.Error())
		}
		sum += value
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		var num1 = -1
		var num2 = -1
		for i := 0; i < len(line) && num1 == -1; i++ {
			num1 = MatchAnyFormatNumeric(line, i)
		}
		for j := len(line) - 1; j >= 0 && num2 == -1; j-- {
			num2 = MatchAnyFormatNumeric(line, j)
		}
		value, _ := strconv.Atoi(strconv.Itoa(num1) + strconv.Itoa(num2))
		sum += value
	}
	return sum
}
func main() {
	lines, err := utils.LoadLinesFromFile("day1.txt")
	if err != nil {
		log.Panic(err.Error())
	}
	log.Printf("Sum for part 1 is %d", part1(lines))
	log.Printf("Sum for part 2 is %d", part2(lines))
}
