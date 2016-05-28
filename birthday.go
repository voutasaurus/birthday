package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	flagHoles   = flag.Int64("h", 0, "specify the number of holes available")
	flagPigeons = flag.Int64("p", 0, "specify the number of pigeons to place randomly in holes")
)

func main() {
	flag.Parse()
	holes := float64(*flagHoles)
	pigeons := float64(*flagPigeons)

	p := prob(holes, pigeons)
	fmt.Printf("%.2f%c\n", 100*p, '%')
}

// prob computes the probability of a collision i:e two pigeons share the same hole.
// The probability is computed using the closed form solution: p = 1 - (holes -1 / holes)^C(pigeons,2).
// Which can be rewritten as: p = 1 - 2^(n)(n-1/2)[log(h -1) - log(h)]
func prob(holes, pigeons float64) float64 {
	l := math.Log2(holes-1) - math.Log2(holes)
	l *= pigeons * (pigeons - 1) / 2
	return 1 - math.Pow(2, l)
}
