package p5

import "testing"

func TestGeneratePrimes_Gen(t *testing.T) {

	if NewPrimes(0).Gen() != nil {
		t.Fatal("0 primes failed!")
	}

	if NewPrimes(1).Gen() != nil {
		t.Fatal("1 primes failed!")
	}

	if len(NewPrimes(5).Gen()) != 3 {
		t.Fatal("5 primes failed!")
	}

	if len(NewPrimes(100).Gen()) != 25 {
		t.Fatal("5 primes failed!")
	}
}
