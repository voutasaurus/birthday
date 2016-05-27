package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
)

var (
	flagHoles   = flag.Int("h", 0, "specify the number of holes available")
	flagPigeons = flag.Int("p", 0, "specify the number of pigeons to place randomly in holes")
)

func main() {
	flag.Parse()
	holes := int64(*flagHoles)
	pigeons := int64(*flagPigeons)

	pCollision, rounding := prob(holes, pigeons)
	fmt.Print(pCollision)
	fmt.Fprintf(os.Stderr, "%s\n", sayRounding[rounding])
}

var sayRounding = map[int]string{
	-1: " (rounded up)",
	0:  " (exactly)",
	1:  " (rounded down)",
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
	probability.SetFrac(numerator.MulRange(m-n+1, m), pow(m, n))

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

	return p, b
}
