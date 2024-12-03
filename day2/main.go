package main

import "fmt"

func main() {
	numSafeReports := CountSafeReports()
	fmt.Println("Part 1: ", numSafeReports)

	numSafeReportsTolerant := TolerantCountSafeReports()
	fmt.Println("Part 2: ", numSafeReportsTolerant)
}
