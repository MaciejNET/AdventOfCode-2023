package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func prefix(word string) int {
	switch {
	case strings.HasPrefix(word, "one"):
		return 1
	case strings.HasPrefix(word, "two"):
		return 2
	case strings.HasPrefix(word, "three"):
		return 3
	case strings.HasPrefix(word, "four"):
		return 4
	case strings.HasPrefix(word, "five"):
		return 5
	case strings.HasPrefix(word, "six"):
		return 6
	case strings.HasPrefix(word, "seven"):
		return 7
	case strings.HasPrefix(word, "eight"):
		return 8
	case strings.HasPrefix(word, "nine"):
		return 9
	default:
		return -1
	}
}

func main() {
	file, err := os.Open("day1.txt");
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		var digits []int
		var num int
		for i, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, int(char - '0'))
			} else {
				if value := prefix(line[i:]); value != -1 {
					digits = append(digits, value)
				}
			}
		}
		length := len(digits)
		if length == 1 {
			num = digits[0] * 10 + digits[0]
		} else {
			num = digits[0] * 10 + digits[length-1]
		}
		sum += num
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Sum:", sum)
}