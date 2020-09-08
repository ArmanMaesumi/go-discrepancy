# go-discrepancy
Go package for approximating star discrepancy, which can be expressed as,

![equation](https://github.com/ArmanMaesumi/go-discrepancy/blob/master/star_discrep_formula.png)

where Q is a rectangular solid in [0, 1]^s, and we have N points (x_1, ..., x_N) in the s-dimensional unit cube.

This code is directly translated from John Burkardt's implementation of "Eric Thiemard, An Algorithm to Compute Bounds for the Star Discrepancy, Journal of Complexity, Volume 17, pages 850-880, 2001."

John Burkardt's implementation: https://people.sc.fsu.edu/~jburkardt/cpp_src/star_discrepancy/star_discrepancy.html
