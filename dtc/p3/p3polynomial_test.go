package p3

import "testing"

func TestPolynomial(t *testing.T) {
	p1 := &polynomial{}
	p1.highPower = 1000
	p1.coeffArray = make([]int, p1.highPower+1)
	p1.coeffArray[1000] = 10
	p1.coeffArray[14] = 5
	p1.coeffArray[0] = 1

	p2 := &polynomial{}
	p2.highPower = 1990
	p2.coeffArray = make([]int, p2.highPower+1)
	p2.coeffArray[1990] = 3
	p2.coeffArray[1492] = -2
	p2.coeffArray[1] = 11
	p2.coeffArray[0] = 5
	PrintPolynomial(AddPolynomial(p1, p2))
	PrintPolynomial(MultPolynomial(p1, p2))
}

func TestNewPolynomialList(t *testing.T) {
	p1 := NewPolynomialList()
	p1.AddList(1000, 10)
	p1.AddList(14, 5)
	p1.AddList(0, 1)
	p2 := NewPolynomialList()
	p2.AddList(1990, 3)
	p2.AddList(1492, -2)
	p2.AddList(1, 11)
	p2.AddList(0, 5)

	p := p1.Add(p2)
	p.print()
	p = p1.mult(p2)
	p.print()
}
