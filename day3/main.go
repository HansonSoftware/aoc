package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Part 1: ", Part1())
	fmt.Println("Part 2: ", Part2())
}

func Part1() int64 {
	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int64 = 0

	for scanner.Scan() {
		result += MultiplyValid(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func MultiplyValid(input string) int64 {
	var result int64 = 0

	// we are looking for mul(X,Y) where X and Y can be between 1-3 digits
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := regex.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		X, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}

		Y, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}

		result += (int64(X) * int64(Y))
	}

	return result
}

func Part2() int64 {
	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int64 = 0
	// mathEnabled := true

	for scanner.Scan() {
		// mul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		// do := regexp.MustCompile(`do\(\)`)
		// dont := regexp.MustCompile(`don't\(\)`)

		// line := scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
