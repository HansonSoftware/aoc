package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func TolerantCountSafeReports() int {
	numSafeReports := 0

	file, err := os.OpenFile("input", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		if TolerantCheckReport(split) == true {
			numSafeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numSafeReports
}

func TolerantCheckReport(report []string) bool {
	var nums []int32

	for i := 0; i < len(report); i++ {
		num, err := strconv.ParseInt(report[i], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, int32(num))
	}

	initialDirection := DetermineDirection(nums[0], nums[1])

	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			break
		}

		currentDirection := DetermineDirection(nums[i], nums[i+1])

		if currentDirection != initialDirection {
			// check if solution would be fine without the problematic level
			// numsWithProblemRemoved := append(nums[:i], nums[i+1:]...)
			// fmt.Println(nums)
			// fmt.Println(numsWithProblemRemoved)
			return false
		}

		diff := Max(nums[i], nums[i+1]) - Min(nums[i], nums[i+1])

		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
