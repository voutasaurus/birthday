package main

import (
	"fmt"
	"math/big"
)

func fac(x int64) *big.Int {
	factorial := new(big.Int)
	return factorial.MulRange(1, x)
}

func pow(x, y int64) *big.Int {
	power := new(big.Int)
	return power.Exp(big.NewInt(x), big.NewInt(y), nil)
}

func prob(m, n int64) (float64, int) {

	probability := new(big.Rat)
	numerator := new(big.Int)

	// Probability that the n items all have their own boxes
	// given m available boxes
	probability.SetFrac(numerator.MulRange(m-n+1, m), pow(m,n))
	
	// Probability that there is some overlap somewhere
	probability.Sub(big.NewRat(1, 1), probability)
	
	// Floating point version (for display)
	p, _ := probability.Float64()
	
	// Check to see whether p is overestimate or underestimate
	// p > probability: b = -1
	// p == probability: b = 0
	// p < probability: b = 1
	z := new(big.Rat)
	b := probability.Cmp(z.SetFloat64(p))
	
	return p,b
}

func main() {
	//fmt.Println(fac(2))
	//fmt.Println(pow(10, 10))
	p, b := prob(100000000000,44500)
	fmt.Println(p, b)
}