package main

import (
	"fmt"
)

func main() {
	score := 85.0 // predefined score

	grade := defineGrade(score)

	fmt.Printf("The grade for the score %.2f is: %s\n", score, grade)
}

func defineGrade(score float64) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}
