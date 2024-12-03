package main

type Direction int

const (
	Increasing Direction = iota
	Decreasing
)

func DetermineDirection(a, b int32) Direction {
	if a < b {
		return Increasing
	}
	return Decreasing
}

func Min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}
