package main

import (
	"fmt"
	"github.com/ArmanMaesumi/go-discrepancy/star"
)

func main()  {
	epsilon := 0.5

	points := [][]float64 {
		{0.1, 0.3},
		{0.2, 0.5},
	}

	lower, upper := star.StarDiscrepancy(epsilon, points, 1, 2)
	fmt.Println(lower)
	fmt.Println(upper)
}