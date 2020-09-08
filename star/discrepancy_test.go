package star

import (
	"testing"
	"math/rand"
)

var lower float64
var upper float64

func benchmarkStarDiscrepancy(n int, s int, b *testing.B) {
	points := make([][]float64, n)
	for i := range points {
		points[i] = make([]float64, s)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < s; j++ {
			points[i][j] = rand.Float64()
		}
	}
	
	var l float64
	var u float64

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l, u = StarDiscrepancy(0.5, points, 1, 2)
	}

	lower = l
	upper = u
}

func BenchmarkStarDiscrepancyDimensions2(b *testing.B) { benchmarkStarDiscrepancy(100, 2, b) }
func BenchmarkStarDiscrepancyDimensions5(b *testing.B) { benchmarkStarDiscrepancy(100, 5, b) }
func BenchmarkStarDiscrepancyDimensions10(b *testing.B) { benchmarkStarDiscrepancy(100, 10, b) }