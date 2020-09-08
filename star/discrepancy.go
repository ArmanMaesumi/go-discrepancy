package star

import (
	"math"
)

type root struct {
	id int
	pt []root
}

var epsilon float64
var n int
var s int
var num int
var den int
var num2 int
var subtotala int = 0
var subtotalb int = 0
var lexsizealpha []int
var maxsizealpha []int
var lexsizebeta []int
var maxsizebeta []int
var borne_sup float64 = 0.0
var borne_inf float64 = 0.0
var points []float64
var suptree []root
var lexalpha [][]root
var lexbeta [][]root

/**
ep - 			0.0 < epsilon < 1.0 (resolution of approximation)
pts - 			matrix of n points in [0, 1]^s
numerator - 	optional balance numerator, use 1 as default
denominator -	optional balance denominator, use 2 as default
				0.0 < numerator / denominator < 1.0

returns:		float64 (lower bound), float64 (upper bound)
**/
func StarDiscrepancy(ep float64, pts [][]float64, numerator int, denominator int) (float64, float64) {
	num = numerator
	den = denominator
	num2 = den - num

	epsilon = ep
	n = len(pts)
	s = len(pts[0])

	points = make([]float64, n * s)
	idx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < s; j++ {
			points[idx] = pts[i][j]
			idx++
		}
	}

	supertree()

	initlex()

	oalpha := make([]float64, s)
	obeta := make([]float64, s)

	for i := 0; i < s; i++ {
		obeta[i] = 1.0
	}

	decomposition(oalpha, obeta, 0, 1.0)

	return borne_inf, borne_sup
}

func decomposition(alpha []float64, beta []float64, min int, value float64) {
	pbetaminp := float64(1.0)
	palpha := float64(1.0)
	var pbeta float64
	var delta float64
	subalpha := make([]float64, s)
	subbeta := make([]float64, s + 1)
	gamma := make([]float64, s + 1)

	for i := min; i < s; i++ {
		pbetaminp *= beta[i];
	}
	pbeta = pbetaminp

	for i := 0; i < min; i++ {
		pbetaminp *= beta[i]
		palpha *= alpha[i]
	}

	pbetaminp -= epsilon
	delta = math.Pow(pbetaminp / (pbeta * palpha), 1.0 / (float64(s) - float64(min)))

	for i := 0; i < min; i++ {
		gamma[i] = alpha[i]
		subalpha[i] = gamma[i]
		subbeta[i] = beta[i]
	}

	for i := min; i < s; i++ {
		gamma[i] = delta * beta[i];
		subalpha[i] = alpha[i];
		subbeta[i] = beta[i];
	}

	subbeta[min] = gamma[min];
	value *= delta;

	if epsilon < value {
		for i := min; i < s; i++ {
			decomposition(subalpha, subbeta, i, value);
			subalpha[i]  = gamma[i];
			subbeta[i]   = beta[i];
			subbeta[i+1] = gamma[i+1];
		}
	} else {
		for i := min; i < s; i++ {
			if i == 0 {
				traiter(subalpha, subbeta, 0);
			} else {
				traiter(subalpha, subbeta, i - 1)
			}
		  
			subalpha[i]  = gamma[i];
			subbeta[i]   = beta[i];
			subbeta[i+1] = gamma[i+1];
		}
	}

	traiter(gamma, beta, s-1);
}

func lowbound(npoints int, volume float64, pave []float64) float64 {
	var tmp float64

	if borne_inf < math.Abs(volume - (float64(npoints) / float64(n))) {
		if volume < (float64(npoints) / float64(n)) {
			volume = 1.0
			for j := 0; j < s; j++ {
				tmp = 0.0
				for i := 0; i < n; i++ {
					if tmp < points[j+i*s] && points[j+i*s] <= pave[j] {
						tmp = points[j+i*s]
					}
				}	
				volume *= tmp			
			}
		} else {
			volume = 1.0
			for j := 0; j < s; j++ {
				tmp = 1.0
				for i := 0; i < n; i++ {
					if points[j+i*s] < tmp && pave[j] <= points[j+i*s] {
						tmp = points[j+i*s]
					}
				}
				volume *= tmp
			}
		}	
		return math.Abs(volume - (float64(npoints) / float64(n)))
	} else {
		return borne_inf
	}
}

func traiter(outputalpha []float64, outputbeta []float64, rng int) {
	valpha := float64(1.0)
	vbeta := float64(1.0)
	var newborn float64
	var nalpha int
	var nbeta int

	for i := 0; i < s; i++ {
		valpha *= outputalpha[i]
		vbeta *= outputbeta[i]
	}

	nalpha = fastexplore(outputalpha, rng, maxsizealpha, lexsizealpha, lexalpha, &subtotala)
	nbeta = fastexplore(outputbeta, rng, maxsizebeta, lexsizebeta, lexbeta, &subtotalb)

	newborn = (float64(nbeta) / float64(n)) - valpha

	if borne_sup < newborn {
		borne_sup = newborn
	}

	newborn = vbeta - (float64(nalpha) / float64(n))

	if borne_sup < newborn {
		borne_sup = newborn
	}

	borne_inf = lowbound(nalpha, valpha, outputalpha)
	borne_inf = lowbound(nbeta, vbeta, outputbeta)
}