# go-discrepancy
Go package for approximating star discrepancy, which can be expressed as,

![equation](https://github.com/ArmanMaesumi/go-discrepancy/blob/master/star_discrep_formula.png)

where Q is a rectangular solid in [0, 1]^s, and we have N points (x_1, ..., x_N) in the s-dimensional unit cube.

This code is directly translated from John Burkardt's implementation of "Eric Thiemard, An Algorithm to Compute Bounds for the Star Discrepancy, Journal of Complexity, Volume 17, pages 850-880, 2001."

John Burkardt's implementation: https://people.sc.fsu.edu/~jburkardt/cpp_src/star_discrepancy/star_discrepancy.html

## Installation
```go get github.com/ArmanMaesumi/go-discrepancy```

## Usage
```
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

          /**
          Parameters:
          ep          - 	0.0 < epsilon < 1.0 (resolution of approximation)
          pts         - 	matrix of n points in [0, 1]^s
          numerator   - 	optional balance numerator, use 1 as default
          denominator -	  optional balance denominator, use 2 as default
                          0.0 < numerator / denominator < 1.0

          returns:        float64 (lower bound), float64 (upper bound)
          **/
	lower, upper := star.StarDiscrepancy(epsilon, points, 1, 2)
	fmt.Println(lower)
	fmt.Println(upper)
}
```
Output:
```
0.9
1
```

## Benchmarks
Random point sets with N=100 are generated with S={2,5,10}. See `star/discrepancy_test.go` for more details
```
BenchmarkStarDiscrepancyDimensions2-16            146623              8585 ns/op
BenchmarkStarDiscrepancyDimensions5-16              5468            213480 ns/op
BenchmarkStarDiscrepancyDimensions10-16               54          19662859 ns/op
```
