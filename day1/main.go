package main

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

// advent of code (aoc) 2024: day 1
func main() {
	scanner1 := ReadFile()
	fmt.Println("Part 1: ", DistanceBetweenLists(scanner1))

	scanner2 := ReadFile()
	fmt.Println("Part 2: ", SimilarityScore(scanner2))
}

/*
* Maybe the lists are only off by a small amount!
* To find out, pair up the numbers and measure how far apart they are.
* Pair up the smallest number in the left list with the smallest number in the right list,
* then the second-smallest left number with the second-smallest right number, and so on.
*
* Within each pair, figure out how far apart the two numbers are;
* you'll need to add up all of those distances.
* For example, if you pair up a 3 from the left list with a 7 from the right list,
* the distance apart is 4; if you pair up a 9 with a 3, the distance apart is 6.
 */
func DistanceBetweenLists(scanner bufio.Scanner) int64 {
	var left []int64
	var right []int64

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		l, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			log.Fatal(err)
			return -1
		}

		r, err := strconv.ParseInt(split[3], 10, 64)
		if err != nil {
			log.Fatal(err)
			return -1
		}

		left = append(left, l)
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(left)
	slices.Sort(right)

	var distance int64 = 0

	for i := 0; i < len(left); i++ {
		max := Max(left[i], right[i])
		min := Min(left[i], right[i])
		difference := max - min
		distance += difference
	}

	return distance
}

/*
* This time, you'll need to figure out exactly how often each number from the left
* list appears in the right list. Calculate a total similarity score by adding up
* each number in the left list after multiplying it by the number of times that
* number appears in the right list.
 */
func SimilarityScore(scanner bufio.Scanner) int64 {
	hashmap := make(map[int64]int64)
	var right []int64

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		l, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			log.Fatal(err)
			return -1
		}

		r, err := strconv.ParseInt(split[3], 10, 64)
		if err != nil {
			log.Fatal(err)
			return -1
		}

		// hashmap [key: left number from file, val: frequency it appears in right]
		hashmap[l] = 0
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(right); i++ {
		val, ok := hashmap[right[i]]
		// if ok is true, key exists
		if ok {
			val += 1
			hashmap[right[i]] = val
		}
	}

	var similarity_score int64 = 0

	for key, val := range hashmap {
		similarity_score += key * val
	}

	return similarity_score
}
