package dtc

import "fmt"

type polynomial struct {
	coeffArray []int
	highPower  int
}

func ZeroPolynomial() *polynomial {
	return &polynomial{}
}

func AddPolynomial(p1, p2 *polynomial) *polynomial {
	pSum := ZeroPolynomial()
	pSum.highPower = p1.highPower
	if p2.highPower > pSum.highPower {
		pSum.highPower = p2.highPower
	}
	pSum.coeffArray = make([]int, pSum.highPower+1)
	for k, v := range p1.coeffArray {
		pSum.coeffArray[k] += v
	}

	for k, v := range p2.coeffArray {
		pSum.coeffArray[k] += v
	}

	return pSum
}

func MultPolynomial(p1, p2 *polynomial) *polynomial {
	multp := ZeroPolynomial()
	multp.highPower = p1.highPower + p2.highPower
	multp.coeffArray = make([]int, multp.highPower+1)
	for k, v := range p1.coeffArray {
		for _k, _v := range p2.coeffArray {

			multp.coeffArray[k+_k] += v * _v
		}
	}
	return multp
}

func PrintPolynomial(p *polynomial) {
	for i := p.highPower; i >= 0; i-- {
		if p.coeffArray[i] != 0 {
			fmt.Printf(" %d*X^%d + ", p.coeffArray[i], i)
		}
	}
	fmt.Println()
}

type polynomialNode struct {
	coefficient int // 系数
	exponent    int // 幂
	next        *polynomialNode
}

func NewPolynomialList() *polynomialNode {
	return &polynomialNode{}
}

func (p *polynomialNode) FindPrevious(exponent int) *polynomialNode {
	if p.next == nil || p.next.exponent <= exponent {
		return p
	}
	tmp := p.next
	for tmp != nil && tmp.next != nil {
		if tmp.exponent > exponent && tmp.next.exponent <= exponent {
			return tmp
		}
		tmp = tmp.next
	}
	return tmp
}

func (p *polynomialNode) AddList(exponent, coefficient int) {
	previous := p.FindPrevious(exponent)
	if previous.next == nil {
		previous.next = &polynomialNode{coefficient: coefficient, exponent: exponent}
	} else {
		if previous.next.exponent == exponent {
			previous.next.coefficient += coefficient
		} else {
			tmp := previous.next
			previous.next = &polynomialNode{coefficient: coefficient, exponent: exponent, next: tmp}
		}
	}
}

func (p *polynomialNode) Add(a *polynomialNode) *polynomialNode {
	sum := NewPolynomialList()
	tmp := sum
	p1 := p.next
	if p1 == nil {
		return a
	}
	p2 := a.next
	if p2 == nil {
		return p
	}
	for p1 != nil && p2 != nil {
		if p1.exponent > p2.exponent {
			tmp.next = &polynomialNode{coefficient: p1.coefficient, exponent: p1.exponent}
			tmp = tmp.next
			p1 = p1.next
		} else if p1.exponent < p2.exponent {
			tmp.next = &polynomialNode{coefficient: p2.coefficient, exponent: p2.exponent}
			tmp = tmp.next
			p2 = p2.next
		} else {
			tmp.next = &polynomialNode{coefficient: p2.coefficient + p1.coefficient, exponent: p2.exponent}
			p1 = p1.next
			p2 = p2.next
		}
	}
	for p1 != nil {
		tmp.next = &polynomialNode{coefficient: p1.coefficient, exponent: p1.exponent}
		tmp = tmp.next
		p1 = p1.next
	}
	for p2 != nil {
		tmp.next = &polynomialNode{coefficient: p2.coefficient, exponent: p2.exponent}
		tmp = tmp.next
		p2 = p2.next
	}
	return sum
}

func (p *polynomialNode) mult(a *polynomialNode) *polynomialNode {
	mult := NewPolynomialList()
	p1 := p.next

	for p1 != nil {
		p2 := a.next
		for p2 != nil {
			mult.AddList(p1.exponent+p2.exponent, p1.coefficient*p2.coefficient)
			p2 = p2.next
		}
		p1 = p1.next
	}
	return mult
}

func (p *polynomialNode) print() {
	tmp := p.next
	fmt.Printf("list: ")
	for tmp != nil {
		fmt.Printf(" %d*X^%d + ", tmp.coefficient, tmp.exponent)
		tmp = tmp.next
	}
	fmt.Println()
}
