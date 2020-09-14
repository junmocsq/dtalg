package p5

import "math"

// 素数计算
type generatePrimes struct {
	crossOut []bool
	maxValue int
}

func NewPrimes(maxValue int) *generatePrimes {
	return &generatePrimes{maxValue: maxValue}
}
func (g *generatePrimes) Gen() []int {
	if g.maxValue >= 2 {
		g.initializeArrayOfBooleans()
		g.crossOutMultiples()
		return g.putUncrossedIntegerIntoResult()
	} else {
		return nil
	}
}

func (g *generatePrimes) initializeArrayOfBooleans() {
	g.crossOut = make([]bool, g.maxValue+1)
	g.crossOut[0] = true
	g.crossOut[1] = true
}

func (g *generatePrimes) crossOutMultiples() {
	maxPrimeFactor := g.calcMaxPrimeFactor()
	for m := 2; m < maxPrimeFactor; m++ {
		if g.notCrossed(m) {
			g.crossOutMultiplesOf(m)
		}
	}
}

func (g *generatePrimes) calcMaxPrimeFactor() int {
	return int(math.Sqrt(float64(g.maxValue))) + 1
}

func (g *generatePrimes) crossOutMultiplesOf(m int) {
	for n := 2 * m; n < g.maxValue+1; n += m {
		g.crossOut[n] = true
	}
}

func (g *generatePrimes) notCrossed(m int) bool {
	return g.crossOut[m] == false
}

func (g *generatePrimes) putUncrossedIntegerIntoResult() []int {
	var result []int
	for m := 2; m < g.maxValue+1; m++ {
		if g.notCrossed(m) {
			result = append(result, m)
		}
	}
	return result
}
