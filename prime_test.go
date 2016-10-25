package prime

import (
	"math/big"
	"testing"
)

var (
	testP            = NewPrimeGenerator(0, 0, nil)
	max2047          = new(big.Int)
	max2048          = new(big.Int)
	knownPrime       = "100010100001111101110100010101011111100001011"
	knownBigPrime    = new(big.Int)
	knownNotBigPrime = new(big.Int)
)

func init() {
	max2047.Exp(big.NewInt(2), big.NewInt(2047), nil)
	max2048.Exp(big.NewInt(2), big.NewInt(2048), nil)
	knownBigPrime.SetString(knownPrime, 2)
	knownNotBigPrime.Mul(knownBigPrime, big.NewInt(2))
}

// Tests

// TestGenerateRandomNumberWithinBitRange tests if the generated number is within the required bit
// range
func TestGenerateRandomNumberWithinBitRange(t *testing.T) {
	n, ok := testP.GenerateRandomNumber()
	if !ok {
		t.Fatal("Error encountered when generating random number")
	}
	lowerBound := n.Cmp(max2047)
	if lowerBound == -1 {
		t.Fatal("Random number less than lower bound")
	}

	upperBound := n.Cmp(max2048)
	if upperBound == 1 {
		t.Fatal("Random number greater than lower bound")
	}
}

// TestProvePrimailtyWithPrime determines the code knows that a known prime is a prime
func TestProvePrimalityWithPrime(t *testing.T) {
	pass := testP.ProvePrimality(knownBigPrime)
	if !pass {
		t.Fatal("Known prime not proved to be prime")
	}
}

// TestProvePrimailityWithNonPrime determines if the code can tell if a known non-prime is a
// non-prime
func TestProvePrimalityWithNonPrime(t *testing.T) {
	pass := testP.ProvePrimality(knownNotBigPrime)
	if pass {
		t.Fatal("Known non-prime proved to be prime")
	}
}

// TestPrimeActuallyGenerated determines if the system actually produces prime numbers
func TestPrimeActuallyGenerated(t *testing.T) {
	p, ok := testP.GetPrime()
	if !ok {
		t.Fatal("Error encountered when generating random prime")
	}

	ok = p.ProbablyPrime(testP.PrimeAccuracy)
	if !ok {
		t.Fatal("Returned a non-prime number")
	}
}

// Benchmarks

// BenchmarkGetPrime benchmarks the overall prime generation process in all of its glory
func BenchmarkGetPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testP.GetPrime()
	}
}

// BenchmarkGenerateRandonNumber benchmarks the large random number generation
func BenchmarkGenerateRandomNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testP.GenerateRandomNumber()
	}
}

// BenchmarkCheckAgainstSmallPrimes benchmarks the initial less expensive check against small known
// primes, with a prime
func BenchmarkCheckAgainstSmallPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testP.checkAgainstSmallPrimes(knownBigPrime)
	}
}

// BenchmarkPrimailityTest benchmarks the whole primaility testing process
func BenchmarkPrimailityTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testP.ProvePrimality(knownBigPrime)
	}
}
